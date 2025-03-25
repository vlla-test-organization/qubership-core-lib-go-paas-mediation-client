package entity

import (
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"testing"
)

func Test_ToConfigMap(t *testing.T) {
	configMapV1 := getConfigMapV1()
	configmap := getEntityConfigMap(configMapV1)
	assert.Equal(t, configMapV1, configmap.ToConfigMap())
}

func Test_NewConfigMap(t *testing.T) {
	v1configmap := getConfigMapV1()
	config := getEntityConfigMap(v1configmap)
	assert.Equal(t, &config, NewConfigMap(v1configmap))
}

func Test_NewConfigMapList(t *testing.T) {
	configMapV1 := getConfigMapV1()
	config := getEntityConfigMap(configMapV1)
	configList := []ConfigMap{config}
	v1configmapList := []*v1.ConfigMap{configMapV1}
	assert.Equal(t, configList, NewConfigMapList(v1configmapList))
}

func getConfigMapV1() *v1.ConfigMap {
	return &v1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: testNamespace, UID: uuid.NewUUID(),
		Annotations: map[string]string{"test": "test1"},
		Labels:      map[string]string{"test": "test2"}},
		Data: map[string]string{"body": "testBody"}}
}

func getEntityConfigMap(v1configmap *v1.ConfigMap) ConfigMap {
	metadata := NewMetadata("ConfigMap", v1configmap.Name, v1configmap.Namespace, string(v1configmap.UID), v1configmap.Generation,
		v1configmap.ResourceVersion, v1configmap.Annotations, v1configmap.Labels)
	return ConfigMap{Metadata: metadata, Data: v1configmap.Data}
}
