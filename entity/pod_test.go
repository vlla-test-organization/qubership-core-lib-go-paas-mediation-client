package entity

import (
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
	"time"
)

func Test_PodStatusFromOsPodStatus(t *testing.T) {
	podCondition := v1.PodCondition{Type: "Ready", Status: "True", LastProbeTime: metav1.Time{Time: time.Now()},
		LastTransitionTime: metav1.Time{Time: time.Now()}}
	podConditionList := []v1.PodCondition{podCondition}
	containerStatus := v1.ContainerStatus{Name: "status", Image: "image", ImageID: "5",
		Ready: true, RestartCount: int32(2), ContainerID: "7",
		State:                v1.ContainerState{Running: &v1.ContainerStateRunning{StartedAt: metav1.Time{Time: time.Now()}}},
		LastTerminationState: v1.ContainerState{Terminated: &v1.ContainerStateTerminated{FinishedAt: metav1.Now()}}}
	containerStatusList := []v1.ContainerStatus{containerStatus}
	podStatusV1 := &v1.PodStatus{PodIP: "1Ip", HostIP: "8080", Phase: "phase", Conditions: podConditionList,
		ContainerStatuses: containerStatusList}

	podStatusTest := PodStatusFromOsPodStatus(podStatusV1)
	assert.Equal(t, 1, len(podStatusTest.ContainerStatuses))
	assert.Equal(t, 1, len(podStatusTest.Conditions))
	assert.Equal(t, "phase", podStatusTest.Phase)
	assert.Equal(t, "1Ip", podStatusTest.PodIP)
	assert.Equal(t, "8080", podStatusTest.HostIP)
}

func Test_PodSpecFromOsPodSpec(t *testing.T) {
	podSpecV1 := v1.PodSpec{NodeName: "podSpec", RestartPolicy: "restart", DNSPolicy: "DNS"}
	mode := int32(2)
	volumes := v1.Volume{Name: "volume1",
		VolumeSource: v1.VolumeSource{Secret: &v1.SecretVolumeSource{SecretName: "secret1", DefaultMode: &mode}}}
	podSpecV1.Volumes = []v1.Volume{volumes}

	container := v1.Container{Name: "1", Image: "nginx", ImagePullPolicy: "pull",
		Resources: v1.ResourceRequirements{Limits: v1.ResourceList{"memory": resource.MustParse("64"),
			"cpu": resource.MustParse("10")},
			Requests: v1.ResourceList{"memory": resource.MustParse("32"), "cpu": resource.MustParse("5")}}}

	volumeMount := v1.VolumeMount{MountPath: "/", Name: "mount", ReadOnly: true}
	container.VolumeMounts = append(container.VolumeMounts, volumeMount)

	containerPort := v1.ContainerPort{ContainerPort: int32(16), Protocol: "protocol", Name: "port"}
	container.Ports = append(container.Ports, containerPort)

	env := v1.EnvVar{Name: "key", Value: "value",
		ValueFrom: &v1.EnvVarSource{FieldRef: &v1.ObjectFieldSelector{APIVersion: "15.1", FieldPath: "/test"},
			SecretKeyRef: &v1.SecretKeySelector{Key: "key",
				LocalObjectReference: v1.LocalObjectReference{Name: "selector"}}}}
	container.Env = append(container.Env, env)

	podSpecV1.Containers = append(podSpecV1.Containers, container)

	podSpecTest := PodSpecFromOsPodSpec(&podSpecV1)
	assert.NotNil(t, podSpecTest)
	assert.Empty(t, podSpecTest.TerminationGracePeriodSeconds)
	assert.Equal(t, 1, len(podSpecTest.Containers))
	assert.Equal(t, 1, len(podSpecTest.Containers[0].VolumeMounts))
	assert.Equal(t, podSpecV1.NodeName, podSpecTest.NodeName)
	assert.Equal(t, podSpecV1.Containers[0].Resources.Limits.Cpu().String(),
		podSpecTest.Containers[0].Resources.Limits.Cpu)
}
