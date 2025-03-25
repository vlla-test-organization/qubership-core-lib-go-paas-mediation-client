package mock

import (
	"context"
	"errors"
	"k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	extensionsv1beta1 "k8s.io/client-go/applyconfigurations/extensions/v1beta1"
)

type FakeIngresses struct {
	Fake *FakeExtensionsV1beta1
	ns   string
}

func (c *FakeIngresses) Create(ctx context.Context, ingress *v1beta1.Ingress, opts v1.CreateOptions) (*v1beta1.Ingress, error) {
	panic("not implemented")
}
func (c *FakeIngresses) Update(ctx context.Context, ingress *v1beta1.Ingress, opts v1.UpdateOptions) (*v1beta1.Ingress, error) {
	panic("not implemented")
}
func (c *FakeIngresses) UpdateStatus(ctx context.Context, ingress *v1beta1.Ingress, opts v1.UpdateOptions) (*v1beta1.Ingress, error) {
	panic("not implemented")
}
func (c *FakeIngresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	panic("not implemented")
}
func (c *FakeIngresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	panic("not implemented")
}
func (c *FakeIngresses) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Ingress, error) {
	return nil, errors.New("error test")
}
func (c *FakeIngresses) List(ctx context.Context, opts v1.ListOptions) (*v1beta1.IngressList, error) {
	return nil, errors.New("WebSocket error test")
}
func (c *FakeIngresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	panic("not implemented")
}
func (c *FakeIngresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Ingress, err error) {
	panic("not implemented")
}
func (c *FakeIngresses) Apply(ctx context.Context, ingress *extensionsv1beta1.IngressApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Ingress, err error) {
	panic("not implemented")
}
func (c *FakeIngresses) ApplyStatus(ctx context.Context, ingress *extensionsv1beta1.IngressApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Ingress, err error) {
	panic("not implemented")
}
