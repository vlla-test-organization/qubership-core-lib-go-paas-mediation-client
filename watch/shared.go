package watch

import (
	"context"

	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type SharedExecutor struct {
}

func (executor *SharedExecutor) CreateWatchRequest(restClient rest.Interface, resource types.PaasResourceType, options *metav1.ListOptions) *rest.Request {
	//TODO implement me
	panic("implement me")
}

func (executor *SharedExecutor) Watch(ctx context.Context, watchRequest *rest.Request) (watch.Interface, error) {
	//TODO implement me
	panic("implement me")
}
