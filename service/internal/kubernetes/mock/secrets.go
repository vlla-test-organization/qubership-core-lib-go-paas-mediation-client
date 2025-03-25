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

type FakeSecrets struct {
	Fake *FakeCoreV1
	ns   string
}

func (c *FakeSecrets) Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) (*v1.Secret, error) {
	panic("not implemented")
}
func (c *FakeSecrets) Update(ctx context.Context, secret *v1.Secret, opts metav1.UpdateOptions) (*v1.Secret, error) {
	panic("not implemented")
}
func (c *FakeSecrets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	panic("not implemented")
}
func (c *FakeSecrets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	panic("not implemented")
}
func (c *FakeSecrets) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Secret, error) {
	return nil, errors.New("error test")
}
func (c *FakeSecrets) List(ctx context.Context, opts metav1.ListOptions) (*v1.SecretList, error) {
	return nil, errors.New("WebSocket error test")
}
func (c *FakeSecrets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	panic("not implemented")
}
func (c *FakeSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Secret, err error) {
	panic("not implemented")
}
func (c *FakeSecrets) Apply(ctx context.Context, secret *corev1.SecretApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Secret, err error) {
	panic("not implemented")
}
