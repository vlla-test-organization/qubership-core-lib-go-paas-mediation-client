package kubernetes

import (
	"context"
	"reflect"
	"sync"
	"testing"
	"time"

	cmv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	certClient "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service/internal/cache"
	"github.com/netcracker/qubership-core-lib-go/v3/logging"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	networkingV1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	fakeWatch "k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

func TestCacheAdapters(t *testing.T) {
	loggerCacheAdapter.SetLevel(logging.LvlDebug)

	for _, cacheType := range []cache.CacheName{
		cache.ConfigMapCache,
		cache.ServiceCache,
		cache.SecretCache,
		cache.RouteCache,
		//cache.CertificateCache, //todo not supported yet
	} {

		assertions := require.New(t)
		ctx := context.Background()

		watchExecutor1 := &testWatcher{channel: make(chan fakeWatch.Event, 5)}
		watchExecutor := &testWatchExecutor{mutex: &sync.Mutex{}, watchers: map[int]watchReturnFunc{
			0: func() (fakeWatch.Interface, error) { return watchExecutor1, nil },
		}}

		clientset := &kubernetes.Clientset{}
		cert_client := &certClient.Clientset{}
		watchHandlers := NewSharedWatchEventHandlers(watchExecutor, watchTimeout,
			clientset.CoreV1().RESTClient(),
			cert_client.CertmanagerV1().RESTClient(),
			clientset.NetworkingV1().RESTClient(),
			clientset.ExtensionsV1beta1().RESTClient())
		resourcesCache := cache.NewTestResourcesCache(cacheType)
		cacheAdapters, err := NewCacheAdapters(ctx, testNamespace1, resourcesCache, watchHandlers)
		assertions.NoError(err)
		assertions.NotNil(cacheAdapters)

		testResourceName := "test-1"

		switch cacheType {
		case cache.ConfigMapCache:
			resourceAdded := createConfigMap(testResourceName, 1, "1", map[string]string{"1": "added"})
			resourceModified := createConfigMap(testResourceName, 2, "2", map[string]string{"1": "modified"})
			resourceDeleted := createConfigMap(testResourceName, 2, "3", map[string]string{"1": "modified"})
			expectedResourceFunc := func(object *v1.ConfigMap) *entity.ConfigMap {
				return entity.NewConfigMap(object)
			}
			resourceCache := resourcesCache.ConfigMaps
			assertions.NotNil(resourceCache)
			testCacheAdapter(ctx, t, testResourceName, resourceAdded, resourceModified, resourceDeleted, expectedResourceFunc, resourceCache, watchExecutor1)
		case cache.ServiceCache:
			resourceAdded := createService(testResourceName, 1, "1")
			resourceModified := createService(testResourceName, 2, "2")
			resourceDeleted := createService(testResourceName, 2, "3")
			expectedResourceFunc := func(object *v1.Service) *entity.Service {
				return entity.NewService(object)
			}
			resourceCache := resourcesCache.Services
			assertions.NotNil(resourceCache)
			testCacheAdapter(ctx, t, testResourceName, resourceAdded, resourceModified, resourceDeleted, expectedResourceFunc, resourceCache, watchExecutor1)
		case cache.SecretCache:
			resourceAdded := createSecret(testResourceName, 1, "1", map[string][]byte{"1": []byte("added")})
			resourceModified := createSecret(testResourceName, 2, "2", map[string][]byte{"1": []byte("modified")})
			resourceDeleted := createSecret(testResourceName, 2, "3", map[string][]byte{"1": []byte("modified")})
			expectedResourceFunc := func(object *v1.Secret) *entity.Secret {
				return entity.NewSecret(object)
			}
			resourceCache := resourcesCache.Secrets
			assertions.NotNil(resourceCache)
			testCacheAdapter(ctx, t, testResourceName, resourceAdded, resourceModified, resourceDeleted, expectedResourceFunc, resourceCache, watchExecutor1)
		case cache.RouteCache:
			resourceAdded := createIngress(testResourceName, 1, "1")
			resourceModified := createIngress(testResourceName, 2, "2")
			resourceDeleted := createIngress(testResourceName, 2, "3")
			expectedResourceFunc := func(object *networkingV1.Ingress) *entity.Route {
				return entity.RouteFromIngressNetworkingV1(object)
			}
			resourceCache := resourcesCache.Ingresses
			assertions.NotNil(resourceCache)
			testCacheAdapter(ctx, t, testResourceName, resourceAdded, resourceModified, resourceDeleted, expectedResourceFunc, resourceCache, watchExecutor1)
		case cache.CertificateCache:
			resourceAdded := createCertificate(testResourceName, 1, "1")
			resourceModified := createCertificate(testResourceName, 2, "2")
			resourceDeleted := createCertificate(testResourceName, 2, "3")
			expectedResourceFunc := func(object *cmv1.Certificate) *entity.Certificate {
				return entity.NewCertificate(object)
			}
			resourceCache := resourcesCache.Certificates
			assertions.NotNil(resourceCache)
			testCacheAdapter(ctx, t, testResourceName, resourceAdded, resourceModified, resourceDeleted, expectedResourceFunc, resourceCache, watchExecutor1)
		}
	}
}

func testCacheAdapter[F runtime.Object, T entity.HasMetadata](
	ctx context.Context,
	t *testing.T,
	resourceName string,
	resourceAdded F,
	resourceModified F,
	resourceDeleted F,
	expectedResourceFunc func(F) *T,
	resourceCache *cache.ResourceCache[T],
	watchExecutor *testWatcher) {
	assertions := require.New(t)

	watchExecutor.channel <- watch.Event{Type: watch.Added, Object: resourceAdded}
	assertions.True(waitUntilPresent(watchTimeout, func() bool {
		resourceFromCache := resourceCache.Get(ctx, testNamespace1, resourceName)
		expectedResource := expectedResourceFunc(resourceAdded)
		return resourceFromCache != nil && reflect.DeepEqual(resourceFromCache, expectedResource)
	}))

	watchExecutor.channel <- watch.Event{Type: watch.Modified, Object: resourceModified}

	assertions.True(waitUntilPresent(watchTimeout, func() bool {
		resourceFromCache := resourceCache.Get(ctx, testNamespace1, resourceName)
		expectedResource := expectedResourceFunc(resourceModified)
		return resourceFromCache != nil && reflect.DeepEqual(resourceFromCache, expectedResource)
	}))

	watchExecutor.channel <- watch.Event{Type: watch.Deleted, Object: resourceDeleted}

	assertions.True(waitUntilPresent(watchTimeout, func() bool {
		resourceFromCache := resourceCache.Get(ctx, testNamespace1, resourceName)
		return resourceFromCache == nil
	}))
}

func waitUntilPresent(timeout time.Duration, targetFunc func() bool) bool {
	timer := time.NewTimer(timeout)
	for {
		select {
		case <-timer.C:
			return false
		default:
			if targetFunc() {
				return true
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
