package entity

import (
	v1 "k8s.io/api/core/v1"
	"time"
)

type (
	Pod struct {
		Metadata `json:"metadata"`
		Spec     PodSpec   `json:"spec"`
		Status   PodStatus `json:"status"`
	}

	PodSpec struct {
		Volumes                       []SpecVolume    `json:"volumes,omitempty"`
		Containers                    []SpecContainer `json:"containers"`
		RestartPolicy                 string          `json:"restartPolicy,omitempty"`
		TerminationGracePeriodSeconds int64           `json:"terminationGracePeriodSeconds,omitempty"`
		DnsPolicy                     string          `json:"dnsPolicy,omitempty"`
		NodeName                      string          `json:"nodeName,omitempty"`
	}

	SpecVolume struct {
		Name   string         `json:"name"`
		Secret *VolumesSecret `json:"secret,omitempty"`
	}

	VolumesSecret struct {
		SecretName  string `json:"secretName,omitempty"`
		DefaultMode int32  `json:"defaultMode,omitempty"`
	}

	SpecContainer struct {
		Name            string                 `json:"name"`
		Image           string                 `json:"image,omitempty"`
		Ports           []ContainerPort        `json:"ports,omitempty"`
		Env             []ContainerEnv         `json:"env,omitempty"`
		Resources       ContainerResources     `json:"resources,omitempty"`
		VolumeMounts    []ContainerVolumeMount `json:"volumeMounts,omitempty"`
		ImagePullPolicy string                 `json:"imagePullPolicy,omitempty"`
		Args            []string               `json:"args"`
	}

	ContainerPort struct {
		ContainerPort int32  `json:"containerPort"`
		Protocol      string `json:"protocol,omitempty"`
		Name          string `json:"name"`
	}

	ContainerEnv struct {
		Name      string     `json:"name"`
		Value     string     `json:"value,omitempty"`
		ValueFrom *ValueFrom `json:"valueFrom,omitempty"`
	}

	ValueFrom struct {
		FieldRef     *FieldRef     `json:"fieldRef,omitempty"`
		SecretKeyRef *SecretKeyRef `json:"secretKeyRef,omitempty"`
	}

	SecretKeyRef struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	}

	FieldRef struct {
		APIVersion string `json:"apiVersion"`
		FieldPath  string `json:"fieldPath"`
	}

	ContainerResources struct {
		Limits   CpuMemoryResource `json:"limits,omitempty"`
		Requests CpuMemoryResource `json:"requests,omitempty"`
	}

	CpuMemoryResource struct {
		Cpu    string `json:"cpu"`
		Memory string `json:"memory"`
	}

	ContainerVolumeMount struct {
		Name      string `json:"name"`
		MountPath string `json:"mountPath"`
		ReadOnly  bool   `json:"readOnly,omitempty"`
	}

	PodStatus struct {
		Phase             string            `json:"phase,omitempty"`
		Conditions        []StatusCondition `json:"conditions,omitempty"`
		HostIP            string            `json:"hostIP,omitempty"`
		PodIP             string            `json:"podIP,omitempty"`
		StartTime         *string           `json:"startTime,omitempty"`
		ContainerStatuses []ContainerStatus `json:"containerStatuses"`
	}

	StatusCondition struct {
		Type               string  `json:"type"`
		Status             string  `json:"status"`
		LastProbeTime      *string `json:"lastProbeTime,omitempty"`
		LastTransitionTime *string `json:"lastTransitionTime,omitempty"`
	}

	ContainerStatus struct {
		Name         string         `json:"name"`
		State        ContainerState `json:"state,omitempty"`
		LastState    ContainerState `json:"lastState,omitempty"`
		Ready        bool           `json:"ready"`
		RestartCount int32          `json:"restartCount"`
		Image        string         `json:"image"`
		ImageID      string         `json:"imageID"`
		ContainerID  string         `json:"containerID,omitempty"`
	}

	ContainerState struct {
		Running    *ContainerStateRunning    `json:"running,omitempty"`
		Terminated *ContainerStateTerminated `json:"terminated,omitempty"`
		Waiting    *ContainerStateWaiting    `json:"waiting,omitempty"`
	}

	ContainerStateRunning struct {
		StartedAt *string `json:"startedAt,omitempty"`
	}

	ContainerStateTerminated struct {
		ContainerID string  `json:"containerID,omitempty"`
		ExitCode    int32   `json:"exitCode"`
		FinishedAt  *string `json:"finishedAt,omitempty"`
		Reason      string  `json:"reason,omitempty"`
		StartedAt   *string `json:"startedAt,omitempty"`
	}

	ContainerStateWaiting struct {
		Message string `json:"message,omitempty"`
		Reason  string `json:"reason,omitempty"`
	}
)

func (pod Pod) GetMetadata() Metadata {
	return pod.Metadata
}

func PodFromOsPod(pod *v1.Pod) *Pod {
	metadata := *FromObjectMeta("Pod", &pod.ObjectMeta)
	return &Pod{Metadata: metadata, Spec: *PodSpecFromOsPodSpec(&pod.Spec), Status: *PodStatusFromOsPodStatus(&pod.Status)}
}

func NewPodList(apiDeploymentList []*v1.Pod) (result []Pod) {
	for _, apiDeployment := range apiDeploymentList {
		deployment := PodFromOsPod(apiDeployment)
		result = append(result, *deployment)
	}
	return result
}

func PodSpecFromOsPodSpec(podSpec *v1.PodSpec) *PodSpec {
	var podVolumesList []SpecVolume
	for _, volume := range podSpec.Volumes {
		var podVolume SpecVolume
		podVolume.Name = volume.Name
		if volume.Secret != nil {
			podVolume.Secret = &VolumesSecret{SecretName: volume.Secret.SecretName, DefaultMode: *volume.Secret.DefaultMode}
		} else {
			podVolume.Secret = nil
		}

		podVolumesList = append(podVolumesList, podVolume)
	}

	var podContainerList []SpecContainer

	for _, container := range podSpec.Containers {
		var podContainer SpecContainer
		podContainer.Name = container.Name
		podContainer.Image = container.Image
		podContainer.ImagePullPolicy = string(container.ImagePullPolicy)

		if container.Resources.Limits != nil {
			podContainer.Resources.Limits.Cpu = container.Resources.Limits.Cpu().String()
			podContainer.Resources.Limits.Memory = container.Resources.Limits.Memory().String()
		}
		if container.Resources.Requests != nil {
			podContainer.Resources.Requests.Cpu = container.Resources.Requests.Cpu().String()
			podContainer.Resources.Requests.Memory = container.Resources.Requests.Memory().String()
		}

		podContainer.Args = container.Args

		for _, volumeMounts := range container.VolumeMounts {
			var containerVolumeMounts ContainerVolumeMount
			containerVolumeMounts.Name = volumeMounts.Name
			containerVolumeMounts.MountPath = volumeMounts.MountPath
			containerVolumeMounts.ReadOnly = volumeMounts.ReadOnly
			podContainer.VolumeMounts = append(podContainer.VolumeMounts, containerVolumeMounts)
		}

		for _, port := range container.Ports {
			var containerPort ContainerPort
			containerPort.ContainerPort = port.ContainerPort
			containerPort.Protocol = string(port.Protocol)
			containerPort.Name = port.Name
			podContainer.Ports = append(podContainer.Ports, containerPort)
		}

		for _, env := range container.Env {
			var containerEnv ContainerEnv
			containerEnv.Name = env.Name
			containerEnv.Value = env.Value

			if vf := env.ValueFrom; vf != nil {
				containerEnv.ValueFrom = NewValueFrom(vf)
			}

			podContainer.Env = append(podContainer.Env, containerEnv)
		}
		podContainerList = append(podContainerList, podContainer)
	}

	var terminationGracePeriodSeconds int64
	if v := podSpec.TerminationGracePeriodSeconds; v != nil {
		terminationGracePeriodSeconds = *v
	}

	return &PodSpec{Volumes: podVolumesList,
		Containers:                    podContainerList,
		RestartPolicy:                 string(podSpec.RestartPolicy),
		TerminationGracePeriodSeconds: terminationGracePeriodSeconds,
		DnsPolicy:                     string(podSpec.DNSPolicy),
		NodeName:                      podSpec.NodeName}
}

func NewValueFrom(vf *v1.EnvVarSource) *ValueFrom {
	if vf.SecretKeyRef != nil {
		return &ValueFrom{SecretKeyRef: &SecretKeyRef{
			Key:  vf.SecretKeyRef.Key,
			Name: vf.SecretKeyRef.Name,
		}}
	}
	if vf.FieldRef != nil {
		return &ValueFrom{FieldRef: &FieldRef{
			APIVersion: vf.FieldRef.APIVersion,
			FieldPath:  vf.FieldRef.FieldPath,
		}}
	}
	return nil
}

func PodStatusFromOsPodStatus(podStatus *v1.PodStatus) *PodStatus {
	var statusConditionsList []StatusCondition

	if podStatus.Conditions != nil {
		for _, condition := range podStatus.Conditions {
			var statusCondition StatusCondition
			statusCondition.Type = string(condition.Type)
			statusCondition.Status = string(condition.Status)

			statusCondition.LastProbeTime = getFormattedTimeString(&condition.LastProbeTime)
			statusCondition.LastTransitionTime = getFormattedTimeString(&condition.LastTransitionTime)

			statusConditionsList = append(statusConditionsList, statusCondition)
		}
	}

	var containerStatusesList []ContainerStatus
	if podStatus.ContainerStatuses != nil {
		for _, containerStatus := range podStatus.ContainerStatuses {
			var entityContainerStatus ContainerStatus

			entityContainerStatus.Name = containerStatus.Name
			entityContainerStatus.Image = containerStatus.Image
			entityContainerStatus.ImageID = containerStatus.ImageID
			entityContainerStatus.Ready = containerStatus.Ready
			entityContainerStatus.RestartCount = containerStatus.RestartCount
			entityContainerStatus.ContainerID = containerStatus.ContainerID
			entityContainerStatus.State = *PodContainerStateFromOsContainerState(&containerStatus.State)
			entityContainerStatus.LastState = *PodContainerStateFromOsContainerState(&containerStatus.LastTerminationState)

			containerStatusesList = append(containerStatusesList, entityContainerStatus)
		}
	}

	startTme := getFormattedTimeString(podStatus.StartTime)

	return &PodStatus{
		Phase:             string(podStatus.Phase),
		Conditions:        statusConditionsList,
		HostIP:            podStatus.HostIP,
		PodIP:             podStatus.PodIP,
		StartTime:         startTme,
		ContainerStatuses: containerStatusesList}
}

func PodContainerStateFromOsContainerState(containerState *v1.ContainerState) *ContainerState {
	var entityContainerState ContainerState

	if containerState.Running != nil {
		if containerState.Running.StartedAt.IsZero() {
			entityContainerState.Running = &ContainerStateRunning{nil}
		} else {
			var timeString = containerState.Running.StartedAt.Format(time.RFC3339)
			entityContainerState.Running = &ContainerStateRunning{&timeString}
		}
	} else {
		entityContainerState.Running = nil
	}

	if containerState.Terminated != nil {
		var startedAt *string
		var finishedAt *string

		startedAt = getFormattedTimeString(&containerState.Terminated.StartedAt)
		finishedAt = getFormattedTimeString(&containerState.Terminated.FinishedAt)

		entityContainerState.Terminated = &ContainerStateTerminated{
			StartedAt:   startedAt,
			FinishedAt:  finishedAt,
			ContainerID: containerState.Terminated.ContainerID,
			ExitCode:    containerState.Terminated.ExitCode,
			Reason:      containerState.Terminated.Reason}

	} else {
		entityContainerState.Terminated = nil
	}

	if containerState.Waiting != nil {
		entityContainerState.Waiting = &ContainerStateWaiting{
			Message: containerState.Waiting.Message,
			Reason:  containerState.Waiting.Reason}
	} else {
		entityContainerState.Waiting = nil
	}

	return &entityContainerState
}
