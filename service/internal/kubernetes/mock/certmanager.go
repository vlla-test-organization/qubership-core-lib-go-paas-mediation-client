package mock

import (
	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned/typed/certmanager/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/testing"
)

type FakeCertmanagerV1 struct {
	*testing.Fake
}

func (f *FakeCertmanagerV1) RESTClient() rest.Interface {
	panic("not implemented")
}

func (f *FakeCertmanagerV1) Certificates(namespace string) certmanagerv1.CertificateInterface {
	return &FakeCertificates{f, namespace}
}

func (f *FakeCertmanagerV1) CertificateRequests(namespace string) certmanagerv1.CertificateRequestInterface {
	panic("not implemented")
}

func (f *FakeCertmanagerV1) ClusterIssuers() certmanagerv1.ClusterIssuerInterface {
	panic("not implemented")
}

func (f *FakeCertmanagerV1) Issuers(namespace string) certmanagerv1.IssuerInterface {
	panic("not implemented")
}
