package entity

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// todo change all implementatios receiver to be a pointer in next major release!
// todo and add SetMetadata(Metadata) method
type HasMetadata interface {
	GetMetadata() Metadata
}

type Metadata struct {
	Kind            string `json:"kind"`
	Name            string `json:"name"`
	Namespace       string `json:"namespace,omitempty"`
	UID             string `json:"uid,omitempty"`
	Generation      int64  `json:"generation,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`

	Annotations map[string]string `json:"annotations,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

func (m *Metadata) ToObjectMeta() *metav1.ObjectMeta {
	if m == nil {
		return nil
	}
	return &metav1.ObjectMeta{
		Name:            m.Name,
		Namespace:       m.Namespace,
		UID:             types.UID(m.UID),
		Generation:      m.Generation,
		ResourceVersion: m.ResourceVersion,
		Annotations:     m.Annotations,
		Labels:          m.Labels,
	}
}

func FromObjectMeta(kind string, m *metav1.ObjectMeta) *Metadata {
	if m == nil {
		return nil
	}
	return &Metadata{
		Kind:            kind,
		Name:            m.Name,
		Namespace:       m.Namespace,
		UID:             string(m.UID),
		Generation:      m.Generation,
		ResourceVersion: m.ResourceVersion,
		Annotations:     m.Annotations,
		Labels:          m.Labels,
	}
}

func NewMetadata(kind, name, namespace, uid string, generation int64, resourceVersion string, annotations, labels map[string]string) Metadata {
	return Metadata{Kind: kind, Name: name, Namespace: namespace, UID: uid, Generation: generation, ResourceVersion: resourceVersion,
		Annotations: annotations, Labels: labels}
}

func NewMetadataFromInterface(kind string, object any) Metadata {
	metadata := Metadata{}
	objectMetadata := object.(map[string]any)
	metadata.Name = objectMetadata["name"].(string)
	if namespace, ok := objectMetadata["namespace"]; ok {
		metadata.Namespace = namespace.(string)
	}
	if generation, ok := objectMetadata["generation"]; ok {
		metadata.Generation = generation.(int64)
	}
	if resourceVersion, ok := objectMetadata["resourceVersion"]; ok {
		metadata.ResourceVersion = resourceVersion.(string)
	}
	if uid, ok := objectMetadata["uid"]; ok {
		metadata.UID = uid.(string)
	}
	metadata.Annotations = ConvertMap(objectMetadata["annotations"])
	metadata.Labels = ConvertMap(objectMetadata["labels"])
	metadata.Kind = kind
	return metadata
}

func ConvertMap(sourceMap any) map[string]string {
	var resultMap = make(map[string]string)
	if sourceMap != nil {
		for key, value := range sourceMap.(map[string]any) {
			strKey := fmt.Sprintf("%v", key)
			strValue := fmt.Sprintf("%v", value)
			resultMap[strKey] = strValue
		}
	}
	return resultMap
}
