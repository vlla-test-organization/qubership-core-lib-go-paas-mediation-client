package kubernetes

import (
	"context"

	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/filter"
	pmWatch "github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/watch"
	corev1 "k8s.io/api/core/v1"
)

func (kube *Kubernetes) GetConfigMap(ctx context.Context, name string, namespace string) (*entity.ConfigMap, error) {
	return GetWrapper(ctx, name, namespace, kube.GetCoreV1Client().ConfigMaps(namespace).Get, kube.Cache.ConfigMaps, entity.NewConfigMap)
}

func (kube *Kubernetes) GetConfigMapList(ctx context.Context, namespace string, filter filter.Meta) ([]entity.ConfigMap, error) {
	return ListWrapper(ctx, filter, kube.GetCoreV1Client().ConfigMaps(namespace).List, kube.Cache.ConfigMaps,
		func(listObj *corev1.ConfigMapList) (result []entity.ConfigMap) {
			for _, item := range listObj.Items {
				result = append(result, *entity.NewConfigMap(&item))
			}
			return
		})
}

func (kube *Kubernetes) WatchConfigMaps(ctx context.Context, namespace string, metaFilter filter.Meta) (*pmWatch.Handler, error) {
	return kube.WatchHandlers.ConfigMaps.Watch(ctx, namespace, metaFilter)
}

func (kube *Kubernetes) DeleteConfigMap(ctx context.Context, name string, namespace string) error {
	return DeleteWrapper(ctx, name, namespace, kube.GetCoreV1Client().ConfigMaps(namespace).Delete)
}

func (kube *Kubernetes) CreateConfigMap(ctx context.Context, configMap *entity.ConfigMap, namespace string) (*entity.ConfigMap, error) {
	return CreateWrapper(ctx, *configMap, kube.GetCoreV1Client().ConfigMaps(namespace).Create, configMap.ToConfigMap, entity.NewConfigMap)
}

func (kube *Kubernetes) UpdateOrCreateConfigMap(ctx context.Context, configMap *entity.ConfigMap, namespace string) (*entity.ConfigMap, error) {
	get := kube.GetCoreV1Client().ConfigMaps(namespace).Get
	create := kube.GetCoreV1Client().ConfigMaps(namespace).Create
	update := kube.GetCoreV1Client().ConfigMaps(namespace).Update
	return UpdateOrCreateWrapper(ctx, *configMap, get, create, update, configMap.ToConfigMap, entity.NewConfigMap)
}
