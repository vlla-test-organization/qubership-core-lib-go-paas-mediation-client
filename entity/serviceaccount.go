package entity

import (
	v1 "k8s.io/api/core/v1"
)

type SecretInfo struct {
	Name string `json:"name"`
}

type ServiceAccount struct {
	Metadata `json:"metadata"`
	Secrets  []SecretInfo `json:"secrets"`
}

func NewServiceAccount(kubeServiceAccount *v1.ServiceAccount) *ServiceAccount {
	metadata := *FromObjectMeta("ServiceAccount", &kubeServiceAccount.ObjectMeta)
	var secrets []SecretInfo
	for _, data := range kubeServiceAccount.Secrets {
		secret := SecretInfo{Name: data.Name}
		secrets = append(secrets, secret)
	}
	return &ServiceAccount{Metadata: metadata, Secrets: secrets}
}

func (sa ServiceAccount) ToServiceAccount() *v1.ServiceAccount {
	var objList []v1.ObjectReference
	for _, data := range sa.Secrets {
		secret := v1.ObjectReference{Name: data.Name}
		objList = append(objList, secret)
	}
	return &v1.ServiceAccount{ObjectMeta: *sa.Metadata.ToObjectMeta(), Secrets: objList}
}

func (sa ServiceAccount) GetMetadata() Metadata {
	return sa.Metadata
}

func NewServiceAccountList(kubeServiceAccountList []*v1.ServiceAccount) []ServiceAccount {
	result := make([]ServiceAccount, 0, len(kubeServiceAccountList))
	for _, kubeServiceAccount := range kubeServiceAccountList {
		result = append(result, *NewServiceAccount(kubeServiceAccount))
	}
	return result
}
