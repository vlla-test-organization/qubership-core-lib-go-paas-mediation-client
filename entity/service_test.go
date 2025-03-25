package entity

import (
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"testing"
)

func getServices() (*v1.Service, *Service) {
	port := Port{Name: "port1",
		Protocol:   "protocol1",
		Port:       int32(4),
		TargetPort: int32(9),
		NodePort:   int32(16)}
	testPorts := []Port{port}
	serviceTest := Service{Metadata: Metadata{Kind: "Service", Name: testService, Namespace: testNamespace},
		Spec: ServiceSpec{Type: "type", ClusterIP: "1", Ports: testPorts}}
	portV1 := v1.ServicePort{Port: port.Port, Name: port.Name, Protocol: v1.Protocol(port.Protocol),
		TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: port.TargetPort}, NodePort: port.NodePort}
	portsV1List := []v1.ServicePort{portV1}
	serviceV1 := v1.Service{ObjectMeta: metav1.ObjectMeta{Name: testService, Namespace: testNamespace},
		Spec: v1.ServiceSpec{Type: "type", ClusterIP: "1", Ports: portsV1List}}
	return &serviceV1, &serviceTest
}
func Test_ToService(t *testing.T) {
	serviceV1, serviceTest := getServices()
	assert.Equal(t, serviceV1, serviceTest.ToService())
}

func Test_NewService(t *testing.T) {
	serviceV1, serviceTest := getServices()
	assert.Equal(t, serviceTest, NewService(serviceV1))
}

func Test_NewServiceList(t *testing.T) {
	serviceV1, serviceTest := getServices()
	serviceV1List := []*v1.Service{serviceV1}
	serviceTestList := []Service{*serviceTest}
	assert.Equal(t, serviceTestList, NewServiceList(serviceV1List))
}
