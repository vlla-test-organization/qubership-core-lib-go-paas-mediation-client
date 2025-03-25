package mock

import (
	"context"
	"errors"
	v1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
)

type FakeCertificates struct {
	Fake *FakeCertmanagerV1
	ns   string
}

func (f *FakeCertificates) Create(ctx context.Context, certificate *v1.Certificate, opts metav1.CreateOptions) (*v1.Certificate, error) {
	panic("not implemented")
}

func (f *FakeCertificates) Update(ctx context.Context, certificate *v1.Certificate, opts metav1.UpdateOptions) (*v1.Certificate, error) {
	panic("not implemented")
}

func (f *FakeCertificates) UpdateStatus(ctx context.Context, certificate *v1.Certificate, opts metav1.UpdateOptions) (*v1.Certificate, error) {
	panic("not implemented")
}

func (f *FakeCertificates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	panic("not implemented")
}

func (f *FakeCertificates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	panic("not implemented")
}

func (f *FakeCertificates) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Certificate, error) {
	return nil, errors.New("error test")
}

func (f *FakeCertificates) List(ctx context.Context, opts metav1.ListOptions) (*v1.CertificateList, error) {
	return nil, errors.New("error test")
}

func (f *FakeCertificates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	panic("not implemented")
}

func (f *FakeCertificates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Certificate, err error) {
	panic("not implemented")
}
