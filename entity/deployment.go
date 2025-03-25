package entity

import (
	v12 "github.com/openshift/api/apps/v1"
	v1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

type Deployment struct {
	Metadata `json:"metadata"`
	Spec     DeploymentSpec   `json:"spec"`
	Status   DeploymentStatus `json:"status"`
}

func (d Deployment) GetMetadata() Metadata {
	return d.Metadata
}

type DeploymentSpec struct {
	Replicas             *int32             `json:"replicas"`
	RevisionHistoryLimit *int32             `json:"revisionHistoryLimit"`
	Template             PodTemplateSpec    `json:"template"`
	Strategy             DeploymentStrategy `json:"strategy"`
}

type DeploymentStrategy struct {
	Type string `json:"type"`
}

type PodTemplateSpec struct {
	Metadata TemplateMetadata `json:"metadata"`
	Spec     PodSpec          `json:"spec"`
}

type TemplateMetadata struct {
	Labels map[string]string `json:"labels,omitempty"`
}

type DeploymentStatus struct {
	AvailableReplicas  int32                 `json:"availableReplicas"`
	Conditions         []DeploymentCondition `json:"conditions"`
	ObservedGeneration int64                 `json:"observedGeneration"`
	ReadyReplicas      int32                 `json:"readyReplicas"`
	Replicas           int32                 `json:"replicas"`
	UpdatedReplicas    int32                 `json:"updatedReplicas"`
}

type DeploymentCondition struct {
	LastTransitionTime *string `json:"lastTransitionTime"`
	LastUpdateTime     *string `json:"lastUpdateTime"`
	Message            string  `json:"message"`
	Reason             string  `json:"reason"`
	Status             string  `json:"status"`
	Type               string  `json:"type"`
}

func NewDeployment(deployment *v1.Deployment) *Deployment {
	metadata := *FromObjectMeta("Deployment", &deployment.ObjectMeta)
	return &Deployment{
		Metadata: metadata,
		Spec:     NewDeploymentSpec(&deployment.Spec),
		Status:   NewDeploymentStatus(&deployment.Status),
	}
}

func NewDeploymentList(apiDeploymentList []*v1.Deployment) []Deployment {
	var result []Deployment
	for _, apiDeployment := range apiDeploymentList {
		deployment := NewDeployment(apiDeployment)
		result = append(result, *deployment)
	}
	return result
}

func NewDeploymentStatus(ds *v1.DeploymentStatus) DeploymentStatus {
	return DeploymentStatus{
		AvailableReplicas:  ds.AvailableReplicas,
		Conditions:         NewConditions(ds.Conditions),
		ObservedGeneration: ds.ObservedGeneration,
		ReadyReplicas:      ds.ReadyReplicas,
		Replicas:           ds.Replicas,
		UpdatedReplicas:    ds.UpdatedReplicas,
	}
}

func NewConditions(dcs []v1.DeploymentCondition) []DeploymentCondition {
	var result []DeploymentCondition
	for _, con := range dcs {
		result = append(result, NewDeploymentCondition(&con))
	}

	return result
}

func NewDeploymentCondition(dc *v1.DeploymentCondition) DeploymentCondition {
	lastTransitionTime := getFormattedTimeString(&dc.LastTransitionTime)
	lastUpdateTime := getFormattedTimeString(&dc.LastUpdateTime)

	return DeploymentCondition{
		LastTransitionTime: lastTransitionTime,
		LastUpdateTime:     lastUpdateTime,
		Message:            dc.Message,
		Reason:             dc.Reason,
		Status:             string(dc.Status),
		Type:               string(dc.Type),
	}
}

func getFormattedTimeString(t *metaV1.Time) *string {
	if t == nil || t.IsZero() {
		return nil
	} else {
		var timeString = t.Format(time.RFC3339)
		return &timeString
	}
}

func NewDeploymentSpec(ds *v1.DeploymentSpec) DeploymentSpec {
	return DeploymentSpec{
		Replicas:             ds.Replicas,
		RevisionHistoryLimit: ds.RevisionHistoryLimit,
		Template:             NewPodTemplateSpec(&ds.Template),
		Strategy:             NewDeploymentStrategy(&ds.Strategy),
	}
}

func NewDeploymentStrategy(ds *v1.DeploymentStrategy) DeploymentStrategy {
	return DeploymentStrategy{
		Type: string(ds.Type),
	}
}

func NewPodTemplateSpec(template *coreV1.PodTemplateSpec) PodTemplateSpec {
	var labels map[string]string
	var spec coreV1.PodSpec
	if template != nil {
		labels = template.Labels
		spec = template.Spec
	}

	return PodTemplateSpec{
		Metadata: NewTemplateMetadata(labels),
		Spec:     *PodSpecFromOsPodSpec(&spec),
	}
}

func NewTemplateMetadata(labels map[string]string) TemplateMetadata {
	return TemplateMetadata{Labels: labels}
}

func NewDeploymentConfig(deployment *v12.DeploymentConfig) *Deployment {
	return &Deployment{
		Metadata: NewMetadata("DeploymentConfig", deployment.Name, deployment.Namespace,
			string(deployment.UID), deployment.Generation, deployment.ResourceVersion,
			deployment.Annotations, deployment.Labels),
		Spec:   NewDeploymentConfigSpec(&deployment.Spec),
		Status: NewDeploymentConfigStatus(&deployment.Status),
	}
}

func NewDeploymentConfigList(apiDeploymentList []*v12.DeploymentConfig) []Deployment {
	var result []Deployment
	for _, apiDeployment := range apiDeploymentList {
		deployment := NewDeploymentConfig(apiDeployment)
		result = append(result, *deployment)
	}
	return result
}

func NewDeploymentConfigSpec(ds *v12.DeploymentConfigSpec) DeploymentSpec {
	return DeploymentSpec{
		Replicas:             &ds.Replicas,
		RevisionHistoryLimit: ds.RevisionHistoryLimit,
		Template:             NewPodTemplateSpec(ds.Template),
		Strategy:             NewDeploymentConfigStrategy(ds.Strategy),
	}
}

func NewDeploymentConfigStrategy(ds v12.DeploymentStrategy) DeploymentStrategy {
	return DeploymentStrategy{
		Type: string(ds.Type),
	}
}

func NewDeploymentConfigStatus(ds *v12.DeploymentConfigStatus) DeploymentStatus {
	return DeploymentStatus{
		AvailableReplicas:  ds.AvailableReplicas,
		Conditions:         NewDeploymentConfigConditions(ds.Conditions),
		ObservedGeneration: ds.ObservedGeneration,
		ReadyReplicas:      ds.ReadyReplicas,
		Replicas:           ds.Replicas,
		UpdatedReplicas:    ds.UpdatedReplicas,
	}
}

func NewDeploymentConfigConditions(dcs []v12.DeploymentCondition) []DeploymentCondition {
	var result []DeploymentCondition
	for _, con := range dcs {
		result = append(result, NewDeploymentConfigCondition(&con))
	}

	return result
}

func NewDeploymentConfigCondition(dc *v12.DeploymentCondition) DeploymentCondition {
	lastTransitionTime := getFormattedTimeString(&dc.LastTransitionTime)
	lastUpdateTime := getFormattedTimeString(&dc.LastUpdateTime)

	return DeploymentCondition{
		LastTransitionTime: lastTransitionTime,
		LastUpdateTime:     lastUpdateTime,
		Message:            dc.Message,
		Reason:             dc.Reason,
		Status:             string(dc.Status),
		Type:               string(dc.Type),
	}
}
