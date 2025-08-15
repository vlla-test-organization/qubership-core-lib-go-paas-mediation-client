package kubernetes

import (
	"context"

	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/filter"
	appv1 "k8s.io/api/apps/v1"
)

func (kube *Kubernetes) GetDeployment(ctx context.Context, resourceName string, namespace string) (*entity.Deployment, error) {
	return GetWrapper(ctx, resourceName, namespace, kube.getAppsV1Client().Deployments(namespace).Get,
		nil, entity.NewDeployment)
}

func (kube *Kubernetes) GetDeploymentList(ctx context.Context, namespace string, filter filter.Meta) ([]entity.Deployment, error) {
	logger.InfoC(ctx, "Search for deployments with filter %v in namespace=%s", filter, namespace)
	return ListWrapper(ctx, filter, kube.getAppsV1Client().Deployments(namespace).List, nil,
		func(listObj *appv1.DeploymentList) (result []entity.Deployment) {
			for _, item := range listObj.Items {
				result = append(result, *entity.NewDeployment(&item))
			}
			return
		})
}
