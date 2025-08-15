package kubernetes

import (
	"context"
	"errors"

	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/types"
	pmWatch "github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/watch"
	paasErrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kube *Kubernetes) GetNamespace(ctx context.Context, name string) (*entity.Namespace, error) {
	ns, err := GetWrapper(ctx, name, name, kube.client.KubernetesInterface.CoreV1().Namespaces().Get, nil, entity.NewNamespace)
	if err != nil && name == kube.namespace {
		return kube.getCurrentNamespaceInternal(), nil
	}
	return ns, err
}

func (kube *Kubernetes) GetNamespaces(ctx context.Context, filter filter.Meta) ([]entity.Namespace, error) {
	namespaceList, err := kube.client.KubernetesInterface.CoreV1().Namespaces().List(ctx, metaV1.ListOptions{})
	if err != nil {
		var statusErr *paasErrors.StatusError
		if ok := errors.As(err, &statusErr); ok && statusErr.Status().Reason == metaV1.StatusReasonForbidden {
			var currentNamespace *entity.Namespace
			logger.WarnC(ctx, "Current service account cannot list namespaces, if you want to list all namespaces create necessary role binding. "+
				"Only current namespace will be returned. Error from kubernetes: %+v", statusErr)
			currentNamespaceFromK8s, errCurrentNS := kube.client.KubernetesInterface.CoreV1().Namespaces().Get(ctx, kube.namespace, metaV1.GetOptions{})
			if errCurrentNS != nil {
				currentNamespace = kube.getCurrentNamespaceInternal()
			} else {
				currentNamespace = entity.NewNamespace(currentNamespaceFromK8s)
			}
			return []entity.Namespace{*currentNamespace}, nil
		}
		return nil, err
	}
	var result []entity.Namespace
	for _, namespace := range namespaceList.Items {
		result = append(result, *entity.NewNamespace(&namespace))
	}
	return result, nil
}

// todo in the next major remove Namespace from meta of Namespace entity
func (kube *Kubernetes) getCurrentNamespaceInternal() *entity.Namespace {
	return &entity.Namespace{Metadata: entity.Metadata{Kind: "Namespace", Name: kube.namespace, Namespace: kube.namespace}}
}

func (kube *Kubernetes) WatchNamespaces(ctx context.Context, namespace string) (*pmWatch.Handler, error) {
	return NewRestWatchHandler(namespace, types.Namespaces, kube.GetCoreV1Client().RESTClient(), kube.WatchExecutor, entity.NewNamespace).
		Watch(ctx, filter.Meta{})
}
