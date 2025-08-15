package kubernetes

import (
	"context"

	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/filter"
	pmWatch "github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/watch"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kube *Kubernetes) CreateServiceAccount(ctx context.Context, serviceAccount *entity.ServiceAccount, namespace string) (*entity.ServiceAccount, error) {
	serviceAccountToCreate := serviceAccount.ToServiceAccount()
	kubeServiceAccount, err := kube.GetCoreV1Client().
		ServiceAccounts(namespace).
		Create(ctx, serviceAccountToCreate, v1.CreateOptions{})
	if err != nil {
		logger.ErrorC(ctx, "Error while creating the service account in kubernetes: %+v", err)
		return nil, err
	}
	return entity.NewServiceAccount(kubeServiceAccount), nil
}

func (kube *Kubernetes) GetServiceAccount(ctx context.Context, resourceName string, namespace string) (*entity.ServiceAccount, error) {
	return GetWrapper(ctx, resourceName, namespace, kube.GetCoreV1Client().ServiceAccounts(namespace).Get,
		nil, entity.NewServiceAccount)
}

func (kube *Kubernetes) GetServiceAccountList(ctx context.Context, namespace string, filter filter.Meta) ([]entity.ServiceAccount, error) {
	return ListWrapper(ctx, filter, kube.GetCoreV1Client().ServiceAccounts(namespace).List, nil,
		func(listObj *corev1.ServiceAccountList) (result []entity.ServiceAccount) {
			for _, item := range listObj.Items {
				result = append(result, *entity.NewServiceAccount(&item))
			}
			return
		})
}

func (kube *Kubernetes) DeleteServiceAccount(ctx context.Context, resourceName string, namespace string) error {
	err := kube.GetCoreV1Client().
		ServiceAccounts(namespace).
		Delete(ctx, resourceName, v1.DeleteOptions{})
	if err != nil {
		logger.ErrorC(ctx, "Error while deleting the service account from kubernetes: %+v", err)
		return err
	}
	return nil
}

func (kube *Kubernetes) WatchServiceAccounts(ctx context.Context, namespace string, metaFilter filter.Meta) (*pmWatch.Handler, error) {
	return kube.WatchHandlers.ServiceAccounts.Watch(ctx, namespace, metaFilter)
}
