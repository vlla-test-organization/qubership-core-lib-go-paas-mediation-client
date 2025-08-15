package kubernetes

import (
	"context"

	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/filter"
)

func (kube *Kubernetes) GetDeploymentFamilyVersions(ctx context.Context, familyName string, namespace string) ([]entity.DeploymentFamilyVersion, error) {
	logger.InfoC(ctx, "Search for deployments with family_name=%s in namespace=%s", familyName, namespace)

	labels := map[string]string{entity.FamilyNameProp: familyName}

	deploymentList, err := kube.GetDeploymentList(ctx, namespace, filter.Meta{Labels: labels})
	if err != nil {
		logger.ErrorC(ctx, "Error getting deployment info: %+v", err)
		return nil, err
	}
	//goland:noinspection GoPreferNilSlice
	result := []entity.DeploymentFamilyVersion{}
	logger.InfoC(ctx, "Found %d resources filtered by family_name=%s", len(deploymentList), familyName)

	for _, deployment := range deploymentList {
		item := entity.DeploymentToDeploymentFamilyVersion(deployment.Metadata.Labels)
		result = append(result, item)
	}
	return result, nil
}
