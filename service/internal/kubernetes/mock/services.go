package mock

import (
	"context"
	"errors"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	corev11 "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/rest"
)

type FakeServices struct {
	Fake *FakeCoreV1
	ns   string
}

func (c *FakeServices) ProxyGet(scheme, name, port, path string, params map[string]string) rest.ResponseWrapper {
	panic("not implemented")
}

func (c *FakeServices) Get(ctx context.Context, name string, opts v1.GetOptions) (*corev1.Service, error) {
	return nil, errors.New("error test")
}

func (c *FakeServices) List(ctx context.Context, opts v1.ListOptions) (*corev1.ServiceList, error) {
	return nil, errors.New("error test")
}

func (c *FakeServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	panic("not implemented")
}

func (c *FakeServices) Create(ctx context.Context, service *corev1.Service, opts v1.CreateOptions) (*corev1.Service, error) {
	panic("not implemented")
}

func (c *FakeServices) Update(ctx context.Context, service *corev1.Service, opts v1.UpdateOptions) (*corev1.Service, error) {
	panic("not implemented")
}

func (c *FakeServices) UpdateStatus(ctx context.Context, service *corev1.Service, opts v1.UpdateOptions) (*corev1.Service, error) {
	panic("not implemented")
}

func (c *FakeServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	panic("not implemented")
}

func (c *FakeServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.Service, err error) {
	panic("not implemented")
}

func (c *FakeServices) Apply(ctx context.Context, service *corev11.ServiceApplyConfiguration, opts v1.ApplyOptions) (result *corev1.Service, err error) {
	panic("not implemented")
}

func (c *FakeServices) ApplyStatus(ctx context.Context, service *corev11.ServiceApplyConfiguration, opts v1.ApplyOptions) (result *corev1.Service, err error) {
	panic("not implemented")
}
