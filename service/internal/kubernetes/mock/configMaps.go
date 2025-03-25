package mock

import (
	"context"
	"errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
)

type FakeConfigMaps struct {
	Fake *FakeCoreV1
	ns   string
}

func (c *FakeConfigMaps) Create(ctx context.Context, configMap *v1.ConfigMap, opts metav1.CreateOptions) (*v1.ConfigMap, error) {
	panic("not implemented")
}
func (c *FakeConfigMaps) Update(ctx context.Context, configMap *v1.ConfigMap, opts metav1.UpdateOptions) (*v1.ConfigMap, error) {
	panic("not implemented")
}
func (c *FakeConfigMaps) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	panic("not implemented")
}
func (c *FakeConfigMaps) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	panic("not implemented")
}
func (c *FakeConfigMaps) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ConfigMap, error) {
	return nil, errors.New("error test")
}
func (c *FakeConfigMaps) List(ctx context.Context, opts metav1.ListOptions) (*v1.ConfigMapList, error) {
	return nil, errors.New("WebSocket error test")
}
func (c *FakeConfigMaps) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	panic("not implemented")
}
func (c *FakeConfigMaps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ConfigMap, err error) {
	panic("not implemented")
}
func (c *FakeConfigMaps) Apply(ctx context.Context, configMap *corev1.ConfigMapApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ConfigMap, err error) {
	panic("not implemented")
}
