package mock

import (
	"k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/testing"
)

type FakeExtensionsV1beta1 struct {
	*testing.Fake
}

func (c *FakeExtensionsV1beta1) DaemonSets(namespace string) v1beta1.DaemonSetInterface {
	panic("not implemented")
}

func (c *FakeExtensionsV1beta1) Deployments(namespace string) v1beta1.DeploymentInterface {
	panic("not implemented")
}

func (c *FakeExtensionsV1beta1) Ingresses(namespace string) v1beta1.IngressInterface {
	return &FakeIngresses{c, namespace}
}

func (c *FakeExtensionsV1beta1) NetworkPolicies(namespace string) v1beta1.NetworkPolicyInterface {
	panic("not implemented")
}

func (c *FakeExtensionsV1beta1) ReplicaSets(namespace string) v1beta1.ReplicaSetInterface {
	panic("not implemented")
}

func (c *FakeExtensionsV1beta1) RESTClient() rest.Interface {
	panic("not implemented")
}
