package mock

import (
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/testing"
)

type FakeCoreV1 struct {
	*testing.Fake
}

func (c *FakeCoreV1) ComponentStatuses() v1.ComponentStatusInterface {
	panic("not implemented")
}

func (c *FakeCoreV1) ConfigMaps(namespace string) v1.ConfigMapInterface {
	return &FakeConfigMaps{c, namespace}
}

func (c *FakeCoreV1) Endpoints(namespace string) v1.EndpointsInterface {
	panic("not implemented")
}

func (c *FakeCoreV1) Events(namespace string) v1.EventInterface {
	panic("not implemented")
}

func (c *FakeCoreV1) LimitRanges(namespace string) v1.LimitRangeInterface {
	panic("not implemented")
}

func (c *FakeCoreV1) Namespaces() v1.NamespaceInterface {
	panic("not implemented")
}

func (c *FakeCoreV1) Nodes() v1.NodeInterface {
	panic("not implemented")
}

func (c *FakeCoreV1) PersistentVolumes() v1.PersistentVolumeInterface {
	panic("not implemented")
}

func (c *FakeCoreV1) PersistentVolumeClaims(namespace string) v1.PersistentVolumeClaimInterface {
	panic("not implemented")
}

func (c *FakeCoreV1) PodTemplates(namespace string) v1.PodTemplateInterface {
	panic("not implemented")
}

func (c *FakeCoreV1) ReplicationControllers(namespace string) v1.ReplicationControllerInterface {
	panic("not implemented")
}

func (c *FakeCoreV1) ResourceQuotas(namespace string) v1.ResourceQuotaInterface {
	panic("not implemented")
}

func (c *FakeCoreV1) Secrets(namespace string) v1.SecretInterface {
	return &FakeSecrets{c, namespace}
}

func (c *FakeCoreV1) Services(namespace string) v1.ServiceInterface {
	return &FakeServices{c, namespace}
}

func (c *FakeCoreV1) Pods(namespace string) v1.PodInterface {
	return &FakePods{c, namespace}
}

func (c *FakeCoreV1) ServiceAccounts(namespace string) v1.ServiceAccountInterface {
	panic("not implemented")
}

func (c *FakeCoreV1) RESTClient() rest.Interface {
	panic("not implemented")
}
