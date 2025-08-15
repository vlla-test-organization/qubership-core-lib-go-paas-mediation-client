package openshiftV3

import (
	"context"

	openshiftappv1 "github.com/openshift/api/apps/v1"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/filter"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/service/internal/kubernetes"
)

func (os *OpenshiftV3Client) GetDeployment(ctx context.Context, resourceName string, namespace string) (*entity.Deployment, error) {
	if result, err := os.Kubernetes.GetDeployment(ctx, resourceName, namespace); err != nil {
		return nil, err
	} else if result != nil {
		return result, nil
	}
	return kubernetes.GetWrapper(ctx, resourceName, namespace, os.AppsClient.DeploymentConfigs(namespace).Get,
		nil, entity.NewDeploymentConfig)
}

func (os *OpenshiftV3Client) GetDeploymentList(ctx context.Context, namespace string, filter filter.Meta) ([]entity.Deployment, error) {
	result, err := os.Kubernetes.GetDeploymentList(ctx, namespace, filter)
	if err != nil {
		return nil, err
	}
	logger.InfoC(ctx, "Search for deploymentConfigs with filter %v in namespace=%s", filter, namespace)
	result2, err := kubernetes.ListWrapper(ctx, filter, os.AppsClient.DeploymentConfigs(namespace).List, nil,
		func(listObj *openshiftappv1.DeploymentConfigList) (result []entity.Deployment) {
			for _, item := range listObj.Items {
				result = append(result, *entity.NewDeploymentConfig(&item))
			}
			return
		})
	if err != nil {
		return nil, err
	}
	return append(result, result2...), nil
}
