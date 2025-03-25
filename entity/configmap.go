package entity

import (
	v1 "k8s.io/api/core/v1"
)

type ConfigMap struct {
	Metadata `json:"metadata"`
	Data     map[string]string `json:"data"`
}

func (m ConfigMap) ToConfigMap() *v1.ConfigMap {
	return &v1.ConfigMap{ObjectMeta: *m.Metadata.ToObjectMeta(), Data: m.Data}
}

func (m ConfigMap) GetMetadata() Metadata {
	return m.Metadata
}

func NewConfigMap(configMap *v1.ConfigMap) *ConfigMap {
	metadata := *FromObjectMeta("ConfigMap", &configMap.ObjectMeta)
	return &ConfigMap{Metadata: metadata, Data: configMap.Data}
}

func NewConfigMapList(kubeConfigMapList []*v1.ConfigMap) []ConfigMap {
	result := make([]ConfigMap, 0)
	for _, kubeConfigMap := range kubeConfigMapList {
		configMap := NewConfigMap(kubeConfigMap)
		result = append(result, *configMap)
	}
	return result
}
