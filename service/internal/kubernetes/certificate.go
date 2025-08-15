package kubernetes

import (
	"context"

	cmv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/filter"
)

func (kube *Kubernetes) GetCertificate(ctx context.Context, resourceName string, namespace string) (*entity.Certificate, error) {
	return GetWrapper(ctx, resourceName, namespace, kube.getCertmanagerV1Client().Certificates(namespace).Get,
		kube.Cache.Certificates, entity.NewCertificate)
}

func (kube *Kubernetes) GetCertificateList(ctx context.Context, namespace string, filter filter.Meta) ([]entity.Certificate, error) {
	return ListWrapper(ctx, filter, kube.getCertmanagerV1Client().Certificates(namespace).List, kube.Cache.Certificates,
		func(listObj *cmv1.CertificateList) (result []entity.Certificate) {
			for _, item := range listObj.Items {
				result = append(result, *entity.NewCertificate(&item))
			}
			return
		})
}

func (kube *Kubernetes) CreateCertificate(ctx context.Context, certificate *entity.Certificate, namespace string) (*entity.Certificate, error) {
	return CreateWrapper(ctx, *certificate, kube.getCertmanagerV1Client().Certificates(namespace).Create,
		certificate.ToCertificate, entity.NewCertificate)
}

func (kube *Kubernetes) UpdateOrCreateCertificate(ctx context.Context, certificate *entity.Certificate, namespace string) (*entity.Certificate, error) {
	get := kube.getCertmanagerV1Client().Certificates(namespace).Get
	create := kube.getCertmanagerV1Client().Certificates(namespace).Create
	update := kube.getCertmanagerV1Client().Certificates(namespace).Update
	return UpdateOrCreateWrapper(ctx, *certificate, get, create, update, certificate.ToCertificate, entity.NewCertificate)
}

func (kube *Kubernetes) DeleteCertificate(ctx context.Context, name string, namespace string) error {
	return DeleteWrapper(ctx, name, namespace, kube.getCertmanagerV1Client().Certificates(namespace).Delete)
}
