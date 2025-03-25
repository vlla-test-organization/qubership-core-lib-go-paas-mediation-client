package entity

import (
	v1 "k8s.io/api/core/v1"
)

type Secret struct {
	Metadata `json:"metadata"`
	Data     map[string][]byte `json:"data"`
	Type     string            `json:"type"`
}

func NewSecret(secret *v1.Secret) *Secret {
	metadata := *FromObjectMeta("Secret", &secret.ObjectMeta)
	data := secret.Data
	secretType := string(secret.Type)
	return &Secret{Metadata: metadata, Data: data, Type: secretType}
}

func NewSecretList(kubeSecretList []*v1.Secret) []Secret {
	result := make([]Secret, 0, len(kubeSecretList))
	for _, kubeSecret := range kubeSecretList {
		result = append(result, *NewSecret(kubeSecret))
	}
	return result
}

func (s Secret) ToSecret() *v1.Secret {
	defer func() {
		if err := recover(); err != nil {
			out, _ := WriteContext()
			logger.Error("panic occurred: %s with secret:%s error:%s", out, s.Name, err)
		} else {
		}
	}()
	return &v1.Secret{ObjectMeta: *s.Metadata.ToObjectMeta(), Data: s.Data, Type: v1.SecretType(s.Type)}
}

func (s Secret) GetMetadata() Metadata {
	return s.Metadata
}
