package kubernetes

import (
	"context"

	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/filter"
	pmWatch "github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/watch"
	corev1 "k8s.io/api/core/v1"
)

func (kube *Kubernetes) GetService(ctx context.Context, resourceName string, namespace string) (*entity.Service, error) {
	return GetWrapper(ctx, resourceName, namespace, kube.GetCoreV1Client().Services(namespace).Get,
		kube.Cache.Services, entity.NewService)
}

func (kube *Kubernetes) DeleteService(ctx context.Context, resourceName string, namespace string) error {
	return DeleteWrapper(ctx, resourceName, namespace, kube.GetCoreV1Client().Services(namespace).Delete)
}

func (kube *Kubernetes) CreateService(ctx context.Context, service *entity.Service, namespace string) (*entity.Service, error) {
	return CreateWrapper(ctx, *service, kube.GetCoreV1Client().Services(namespace).Create, service.ToService, entity.NewService)
}

func (kube *Kubernetes) UpdateOrCreateService(ctx context.Context, service *entity.Service, namespace string) (*entity.Service, error) {
	get := kube.GetCoreV1Client().Services(namespace).Get
	create := kube.GetCoreV1Client().Services(namespace).Create
	update := kube.GetCoreV1Client().Services(namespace).Update
	return UpdateOrCreateWrapper(ctx, *service, get, create, update, service.ToService, entity.NewService)
}

func (kube *Kubernetes) GetServiceList(ctx context.Context, namespace string, filter filter.Meta) ([]entity.Service, error) {
	return ListWrapper(ctx, filter, kube.GetCoreV1Client().Services(namespace).List, kube.Cache.Services,
		func(listObj *corev1.ServiceList) (result []entity.Service) {
			for _, item := range listObj.Items {
				result = append(result, *entity.NewService(&item))
			}
			return
		})
}

func (kube *Kubernetes) WatchServices(ctx context.Context, namespace string, metaFilter filter.Meta) (*pmWatch.Handler, error) {
	return kube.WatchHandlers.Services.Watch(ctx, namespace, metaFilter)
}
