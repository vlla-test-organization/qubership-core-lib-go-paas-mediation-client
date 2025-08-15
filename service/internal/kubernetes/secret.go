package kubernetes

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/entity"
	"github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/filter"
	pmWatch "github.com/vlla-test-organization/qubership-core-lib-go-paas-mediation-client/v8/watch"
	"golang.org/x/mod/semver"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const ServiceAccountNameAnnotation = "kubernetes.io/service-account.name"
const SecretTypeServiceAccountTokenAnnotation = "kubernetes.io/service-account-token"

func (kube *Kubernetes) GetSecret(ctx context.Context, resourceName string, namespace string) (*entity.Secret, error) {
	return GetWrapper(ctx, resourceName, namespace, kube.GetCoreV1Client().Secrets(namespace).Get,
		kube.Cache.Secrets, entity.NewSecret)
}

func (kube *Kubernetes) GetSecretList(ctx context.Context, namespace string, filter filter.Meta) ([]entity.Secret, error) {
	return ListWrapper(ctx, filter, kube.GetCoreV1Client().Secrets(namespace).List, kube.Cache.Secrets,
		func(listObj *corev1.SecretList) (result []entity.Secret) {
			for _, item := range listObj.Items {
				result = append(result, *entity.NewSecret(&item))
			}
			return
		})
}

func (kube *Kubernetes) WatchSecrets(ctx context.Context, namespace string, metaFilter filter.Meta) (*pmWatch.Handler, error) {
	return kube.WatchHandlers.Secrets.Watch(ctx, namespace, metaFilter)
}

func (kube *Kubernetes) CreateSecret(ctx context.Context, secret *entity.Secret, namespace string) (*entity.Secret, error) {
	return CreateWrapper(ctx, *secret, kube.GetCoreV1Client().Secrets(namespace).Create, secret.ToSecret, entity.NewSecret)
}

func (kube *Kubernetes) DeleteSecret(ctx context.Context, resourceName string, namespace string) error {
	return DeleteWrapper(ctx, resourceName, namespace, kube.GetCoreV1Client().Secrets(namespace).Delete)
}

func (kube *Kubernetes) UpdateOrCreateSecret(ctx context.Context, secret *entity.Secret, namespace string) (*entity.Secret, error) {
	get := kube.GetCoreV1Client().Secrets(namespace).Get
	create := kube.GetCoreV1Client().Secrets(namespace).Create
	update := kube.GetCoreV1Client().Secrets(namespace).Update
	return UpdateOrCreateWrapper(ctx, *secret, get, create, update, secret.ToSecret, entity.NewSecret)
}

func (kube *Kubernetes) GetOrCreateSecretWithServiceAccountToken(ctx context.Context, serviceAccountName string, namespace string) (*entity.Secret, error) {
	kubernetesVersion, err := kube.GetKubernetesVersion()
	if err != nil {
		return nil, err
	}
	searchTokenSecretDirectly := semver.Compare(kubernetesVersion, "v1.24") >= 0
	if searchTokenSecretDirectly {
		logger.DebugC(ctx, "searching secret with token directly")
		// check if secret with token was created manually for this sa (kubernetes starting from 1.24 does not create secret with token automatically)
		annotations := map[string]string{ServiceAccountNameAnnotation: serviceAccountName}
		secretList, errL := kube.GetSecretList(ctx, namespace, filter.Meta{
			Annotations: annotations,
		})
		if errL != nil {
			return nil, errL
		}
		if len(secretList) == 1 {
			return &secretList[0], nil
		} else if len(secretList) == 0 {
			// no token found, create it manually
			secret := &corev1.Secret{
				ObjectMeta: v1.ObjectMeta{Name: serviceAccountName + "-token", Namespace: namespace,
					Annotations: map[string]string{ServiceAccountNameAnnotation: serviceAccountName}},
				Type: SecretTypeServiceAccountTokenAnnotation,
			}
			logger.InfoC(ctx, "creating secret with service account token, secret name = %s", secret.Name)
			if createdSecret, errC := kube.UpdateOrCreateSecret(ctx, entity.NewSecret(secret), namespace); errC != nil {
				return nil, errC
			} else {
				logger.InfoC(ctx, "created secret with service account token, secret name = %s", createdSecret.Name)
				attempts := 30
				for attempts > 0 {
					createdSecret, err = kube.GetSecret(ctx, createdSecret.Name, namespace)
					if err != nil {
						return nil, err
					}
					if createdSecret.Data != nil && createdSecret.Data["token"] != nil {
						break
					}
					attempts--
					time.Sleep(2 * time.Second)
				}
				if createdSecret.Data != nil && createdSecret.Data["token"] != nil {
					return createdSecret, nil
				} else {
					return nil, fmt.Errorf("failed to wait for secret '%s' to contain non empty token", createdSecret.Name)
				}
			}
		} else {
			return nil, fmt.Errorf("failed to get service account token secret. multiple secrets found")
		}
	} else {
		logger.DebugC(ctx, "searching secret with token via service account")
		attempts := 30
		var secretName string
		for attempts > 0 {
			sa, err := kube.GetServiceAccount(ctx, serviceAccountName, namespace)
			if err != nil {
				return nil, err
			}
			secretName = getTokenSecretName(serviceAccountName, sa.Secrets)
			if secretName != "" {
				break
			}
			attempts--
			time.Sleep(2 * time.Second)
		}
		if secretName == "" {
			// no secret found in service account
			return nil, fmt.Errorf("failed to wait for service account '%s' to contain secret with token", serviceAccountName)
		} else {
			secret, err := kube.GetSecret(ctx, secretName, namespace)
			if err != nil {
				return nil, err
			}
			return secret, nil
		}
	}
}

func getTokenSecretName(serviceAccount string, secrets []entity.SecretInfo) string {
	for _, secret := range secrets {
		if strings.Contains(secret.Name, serviceAccount+"-token") {
			return secret.Name
		}
	}
	return ""
}
