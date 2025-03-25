package openshiftV3

import (
	kube "github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service/internal/kubernetes"
	"github.com/netcracker/qubership-core-lib-go/v3/logging"
	projectv1 "github.com/openshift/api/project/v1"
	appsv1client "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	projectv1client "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	routeScheme "github.com/openshift/client-go/route/clientset/versioned/scheme"
	routev1client "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func init() {
	routeScheme.AddToScheme(scheme.Scheme)
	if err := projectv1.AddToScheme(scheme.Scheme); err != nil {
		panic(err)
	}
}

var (
	logger        logging.Logger
	resourceAlias = map[string]string{
		"namespaces": "projects",
	}
)

func init() {
	logger = logging.GetLogger("openshift_service3")
}

type OpenshiftV3Client struct {
	*kube.Kubernetes
	RouteV1Client   routev1client.RouteV1Interface
	ProjectV1Client projectv1client.ProjectV1Interface
	AppsClient      appsv1client.AppsV1Interface
}

// We use isLocal flag for convenient development and fast bug fix. I don't recommend to remove this option, because it does no harm.
func GetKubeConfig(isLocal bool) (*rest.Config, error) {
	if isLocal {
		logger.Debug("Run in outside mode and get kube config from home directory")
		clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			clientcmd.NewDefaultClientConfigLoadingRules(),
			&clientcmd.ConfigOverrides{},
		)
		return clientConfig.ClientConfig()
	}
	logger.Debug("Run in inside pod mode")
	return rest.InClusterConfig()
}

func NewOpenshiftV3Client(routeV1Client routev1client.RouteV1Interface,
	projectV1Client projectv1client.ProjectV1Interface,
	appsV1Client appsv1client.AppsV1Interface,
	kubeClient *kube.Kubernetes) *OpenshiftV3Client {
	openshiftClient := OpenshiftV3Client{}
	openshiftClient.RouteV1Client = routeV1Client
	openshiftClient.ProjectV1Client = projectV1Client
	openshiftClient.AppsClient = appsV1Client
	openshiftClient.Kubernetes = kubeClient
	return &openshiftClient
}
