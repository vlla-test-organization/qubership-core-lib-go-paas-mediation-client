package entity

import (
	osV1 "github.com/openshift/api/project/v1"
	v1 "k8s.io/api/core/v1"
)

type Namespace struct {
	Metadata `json:"metadata"`
}

func NewNamespace(namespace *v1.Namespace) *Namespace {
	metadata := *FromObjectMeta("Namespace", &namespace.ObjectMeta)
	metadata.Namespace = namespace.Name
	return &Namespace{Metadata: metadata}
}

func NewNamespaceFromOsProject(project *osV1.Project) *Namespace {
	metadata := *FromObjectMeta("Namespace", &project.ObjectMeta)
	metadata.Namespace = project.Name
	return &Namespace{Metadata: metadata}
}

func NewNamespaceFromInterface(object any) *Namespace {
	metadataObj := object.(map[string]any)["metadata"]
	metadata := NewMetadataFromInterface("Namespace", metadataObj)
	metadata.Namespace = metadata.Name
	return &Namespace{Metadata: metadata}
}

func (n Namespace) GetMetadata() Metadata {
	return n.Metadata
}
