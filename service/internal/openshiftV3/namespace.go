package openshiftV3

import (
	"context"

	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (os *OpenshiftV3Client) GetNamespaces(ctx context.Context, filter filter.Meta) ([]entity.Namespace, error) {
	projectList, e := os.ProjectV1Client.Projects().List(ctx, v1.ListOptions{})
	if e != nil {
		return nil, e
	}
	var result []entity.Namespace
	for _, project := range projectList.Items {
		result = append(result, *entity.NewNamespaceFromOsProject(&project))
	}
	return result, nil
}
