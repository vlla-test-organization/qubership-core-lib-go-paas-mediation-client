package kubernetes

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service/internal/cache"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_WithErrorHandler_WithCache_WithErrFromApiServer(t *testing.T) {
	assertions := require.New(t)
	resourceName := "test-name"
	namespace := "test-namespace"

	getFromApiServer := func(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.ConfigMap, error) {
		return nil, fmt.Errorf("test error from api server")
	}
	cachedConfigMap := entity.ConfigMap{Metadata: entity.Metadata{
		Name:      resourceName,
		Namespace: namespace,
	}}

	resourcesCache, err := cache.NewResourcesCache(10, 1000, 1000, 0)
	assertions.NoError(err)
	ctx := context.Background()
	ok, err := resourcesCache.ConfigMaps.Set(ctx, cachedConfigMap)
	assertions.NoError(err)
	assertions.True(ok)

	result, err := GetWrapper(context.Background(), resourceName, namespace, getFromApiServer, resourcesCache.ConfigMaps, entity.NewConfigMap)
	assertions.NoError(err)
	assertions.NotNil(result)
	assertions.Equal(cachedConfigMap, *result)
}

func Test_WithErrorHandler_WithEmptyCache_WithErrFromApiServer(t *testing.T) {
	assertions := require.New(t)
	resourceName := "test-name"
	namespace := "test-namespace"

	apiErr := fmt.Errorf("test error from api server")
	getFromApiServer := func(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.ConfigMap, error) {
		return nil, apiErr
	}
	resourcesCache, err := cache.NewResourcesCache(10, 1000, 1000, 0)
	assertions.NoError(err)

	result, err := GetWrapper(context.Background(), resourceName, namespace, getFromApiServer, resourcesCache.ConfigMaps, entity.NewConfigMap)
	assertions.Equal(apiErr, err)
	assertions.Nil(result)
}

func Test_WithErrorHandler_WithEmptyCache_WithSuccessFromApiServer(t *testing.T) {
	assertions := require.New(t)
	resourceName := "test-name"
	namespace := "test-namespace"

	configMapFromApiServer := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: resourceName}}
	getFromApiServer := func(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.ConfigMap, error) {
		return configMapFromApiServer, nil
	}

	resourcesCache, err := cache.NewResourcesCache(10, 1000, 1000, 0)
	assertions.NoError(err)

	result, err := GetWrapper(context.Background(), resourceName, namespace, getFromApiServer, resourcesCache.ConfigMaps, entity.NewConfigMap)
	assertions.Nil(err)
	assertions.NotNil(result)
}

func Test_WithErrorHandler_WithoutCache_WithSuccessFromApiServer(t *testing.T) {
	assertions := require.New(t)
	resourceName := "test-name"
	namespace := "test-namespace"

	getFromApiServer := func(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.ConfigMap, error) {
		return &corev1.ConfigMap{}, nil
	}
	result, err := GetWrapper(context.Background(), resourceName, namespace, getFromApiServer, nil, entity.NewConfigMap)
	assertions.Nil(err)
	assertions.NotNil(result)
}

func Test_WithErrorHandler_WithoutCache_WithErrorFromApiServer(t *testing.T) {
	assertions := require.New(t)
	resourceName := "test-name"
	namespace := "test-namespace"

	apiErr := fmt.Errorf("test error from api server")
	getFromApiServer := func(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.ConfigMap, error) {
		return nil, apiErr
	}
	result, err := GetWrapper(context.Background(), resourceName, namespace, getFromApiServer, nil, entity.NewConfigMap)
	assertions.Equal(apiErr, err)
	assertions.Nil(result)
}

func TestContainsAllAnnotations(t *testing.T) {
	assertions := require.New(t)
	var annotationMap = make(map[string]string)
	annotationMap["a1"] = "v1"
	annotationMap["a2"] = "v2"

	var requiredAnns = make(map[string]string)
	requiredAnns["a1"] = "v1"
	requiredAnns["a2"] = "*"
	assertions.True(containsAll(annotationMap, requiredAnns), "expect all annotation are contained")

	requiredAnns["a2"] = "v3"
	assertions.False(containsAll(annotationMap, requiredAnns), "expect annotation are not contained")
}

func TestListWrapper2PartsRequestSuccess(t *testing.T) {
	assertions := require.New(t)
	name1 := "test-1"
	name2 := "test-2"
	name3 := "test-3"
	name4 := "test-4"
	listFromApiServer := func(ctx context.Context, opts metav1.ListOptions) (*corev1.ConfigMapList, error) {
		switch opts.Continue {
		case "":
			return &corev1.ConfigMapList{
				ListMeta: metav1.ListMeta{Continue: "1"},
				Items: []corev1.ConfigMap{
					*createCoreV1ConfigMap(testNamespace1, name1, nil),
					*createCoreV1ConfigMap(testNamespace1, name2, nil),
				},
			}, nil
		case "1":
			return &corev1.ConfigMapList{
				ListMeta: metav1.ListMeta{Continue: ""},
				Items: []corev1.ConfigMap{
					*createCoreV1ConfigMap(testNamespace1, name3, nil),
					*createCoreV1ConfigMap(testNamespace1, name4, nil),
				},
			}, nil
		default:
			return nil, errors.New("invalid continue value")
		}
	}
	resourcesCache, err := cache.NewResourcesCache(2, 1000, 1000, 0, cache.ConfigMapCache)
	assertions.NoError(err)
	ctx := context.Background()
	result, err := ListWrapper(ctx, filter.Meta{}, listFromApiServer, resourcesCache.ConfigMaps,
		func(listObj *corev1.ConfigMapList) (result []entity.ConfigMap) {
			for _, item := range listObj.Items {
				result = append(result, *entity.NewConfigMap(&item))
			}
			return
		})
	assertions.NoError(err)
	assertions.NotNil(result)
	assertions.Equal(4, len(result))
	namesSet := map[string]struct{}{}
	for _, r := range result {
		namesSet[r.Name] = struct{}{}
	}
	assertions.Equal(4, len(namesSet))
}

func TestListWrapper2PartsRequest410GoneResponse(t *testing.T) {
	assertions := require.New(t)
	counter := 0
	name1 := "test-1"
	name2 := "test-2"
	name3 := "test-3"
	name4 := "test-4"
	listFromApiServer := func(ctx context.Context, opts metav1.ListOptions) (*corev1.ConfigMapList, error) {
		switch opts.Continue {
		case "":
			counter++
			return &corev1.ConfigMapList{
				ListMeta: metav1.ListMeta{Continue: strconv.Itoa(counter)},
				Items: []corev1.ConfigMap{
					*createCoreV1ConfigMap(testNamespace1, name1, nil),
					*createCoreV1ConfigMap(testNamespace1, name2, nil),
				},
			}, nil
		case "1":
			return nil, k8sErrors.NewResourceExpired("test expired error")
		case "2":
			return &corev1.ConfigMapList{
				ListMeta: metav1.ListMeta{Continue: ""},
				Items: []corev1.ConfigMap{
					*createCoreV1ConfigMap(testNamespace1, name3, nil),
					*createCoreV1ConfigMap(testNamespace1, name4, nil),
				},
			}, nil
		default:
			return nil, errors.New("invalid continue value")
		}
	}
	resourcesCache, err := cache.NewResourcesCache(2, 1000, 1000, 0, cache.ConfigMapCache)
	assertions.NoError(err)
	ctx := context.Background()
	result, err := ListWrapper(ctx, filter.Meta{}, listFromApiServer, resourcesCache.ConfigMaps,
		func(listObj *corev1.ConfigMapList) (result []entity.ConfigMap) {
			for _, item := range listObj.Items {
				result = append(result, *entity.NewConfigMap(&item))
			}
			return
		})
	assertions.NoError(err)
	assertions.NotNil(result)
	assertions.Equal(4, len(result))
	namesSet := map[string]struct{}{}
	for _, r := range result {
		namesSet[r.Name] = struct{}{}
	}
	assertions.Equal(4, len(namesSet))
}

func TestListWrapper2PartsRequestSuccessWithAnnotationsFilter(t *testing.T) {
	assertions := require.New(t)
	name1 := "test-1"
	name2 := "test-2"
	name3 := "test-3"
	name4 := "test-4"
	listFromApiServer := func(ctx context.Context, opts metav1.ListOptions) (*corev1.ConfigMapList, error) {
		switch opts.Continue {
		case "":
			return &corev1.ConfigMapList{
				ListMeta: metav1.ListMeta{Continue: "1"},
				Items: []corev1.ConfigMap{
					*createCoreV1ConfigMap(testNamespace1, name1, map[string]string{"1": "1"}),
					*createCoreV1ConfigMap(testNamespace1, name2, map[string]string{"2": "2"}),
				},
			}, nil
		case "1":
			return &corev1.ConfigMapList{
				ListMeta: metav1.ListMeta{Continue: ""},
				Items: []corev1.ConfigMap{
					*createCoreV1ConfigMap(testNamespace1, name3, map[string]string{"1": "1"}),
					*createCoreV1ConfigMap(testNamespace1, name4, map[string]string{"2": "2"}),
				},
			}, nil
		default:
			return nil, errors.New("invalid continue value")
		}
	}
	resourcesCache, err := cache.NewResourcesCache(2, 1000, 1000, 0)
	assertions.NoError(err)
	ctx := context.Background()
	result, err := ListWrapper(ctx, filter.Meta{Annotations: map[string]string{"2": "2"}}, listFromApiServer, resourcesCache.ConfigMaps,
		func(listObj *corev1.ConfigMapList) (result []entity.ConfigMap) {
			for _, item := range listObj.Items {
				result = append(result, *entity.NewConfigMap(&item))
			}
			return
		})
	assertions.NoError(err)
	assertions.NotNil(result)
	assertions.Equal(2, len(result))
	namesSet := map[string]struct{}{}
	for _, r := range result {
		namesSet[r.Name] = struct{}{}
	}
	assertions.Equal(2, len(namesSet))
}

func createCoreV1ConfigMap(namespace, name string, annotations map[string]string) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Namespace:   namespace,
			Generation:  1,
			Annotations: annotations,
		},
	}
}
