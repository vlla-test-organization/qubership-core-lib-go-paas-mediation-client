package entity

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type Port struct {
	Name       string `json:"name"`
	Protocol   string `json:"protocol"`
	Port       int32  `json:"port"`
	TargetPort int32  `json:"targetPort"`
	NodePort   int32  `json:"nodePort,omitempty"`
}

type ServiceSpec struct {
	Ports     []Port            `json:"ports"`
	Selector  map[string]string `json:"selector"`
	ClusterIP string            `json:"clusterIP"`
	Type      string            `json:"type"`
}

type Service struct {
	Metadata `json:"metadata"`
	Spec     ServiceSpec `json:"spec"`
}

func NewService(kubeService *v1.Service) *Service {
	metadata := *FromObjectMeta("Service", &kubeService.ObjectMeta)
	var ports []Port
	for _, port := range kubeService.Spec.Ports {
		servicePort := Port{Port: port.Port, Name: port.Name, Protocol: string(port.Protocol), TargetPort: port.TargetPort.IntVal}
		if port.NodePort != 0 {
			servicePort.NodePort = port.NodePort
		}
		ports = append(ports, servicePort)
	}
	serviceSpec := ServiceSpec{}
	serviceSpec.Ports = ports
	serviceSpec.ClusterIP = kubeService.Spec.ClusterIP
	serviceSpec.Selector = kubeService.Spec.Selector
	serviceSpec.Type = string(kubeService.Spec.Type)
	return &Service{Metadata: metadata, Spec: serviceSpec}
}

func (s Service) ToService() *v1.Service {
	serviceSpec := v1.ServiceSpec{}
	serviceSpec.Type = v1.ServiceType(s.Spec.Type)
	serviceSpec.ClusterIP = s.Spec.ClusterIP
	serviceSpec.Selector = s.Spec.Selector
	for _, port := range s.Spec.Ports {
		servicePort := v1.ServicePort{Port: port.Port, Name: port.Name, Protocol: v1.Protocol(port.Protocol),
			TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: port.TargetPort}}
		if port.NodePort != 0 {
			servicePort.NodePort = port.NodePort
		}
		serviceSpec.Ports = append(serviceSpec.Ports, servicePort)
	}
	return &v1.Service{Spec: serviceSpec, ObjectMeta: *s.Metadata.ToObjectMeta()}
}

func (s Service) GetMetadata() Metadata {
	return s.Metadata
}

func NewServiceList(kubeServiceList []*v1.Service) []Service {
	result := make([]Service, 0)
	for _, kubeService := range kubeServiceList {
		service := NewService(kubeService)
		result = append(result, *service)
	}
	return result
}
