package openshiftV3

import (
	"context"

	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/service/internal/kubernetes"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/types"
	pmWatch "github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/watch"
)

func (os *OpenshiftV3Client) WatchNamespaces(ctx context.Context, namespace string) (*pmWatch.Handler, error) {
	return kubernetes.NewRestWatchHandler(namespace, types.Projects, os.ProjectV1Client.RESTClient(), os.Kubernetes.WatchExecutor, entity.NewNamespaceFromOsProject).
		Watch(ctx, filter.Meta{})
}
