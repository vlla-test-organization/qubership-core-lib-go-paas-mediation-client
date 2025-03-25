package entity

import (
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func getSecretV1() *v1.Secret {
	return &v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: testNamespace,
		Annotations: map[string]string{"test": "test1"},
		Labels:      map[string]string{"test": "test2"}}, Data: map[string][]byte{"body": {11, 10, 42, 255}}}
}

func Test_NewSecret(t *testing.T) {
	v1secret := getSecretV1()
	secret := Secret{Metadata: Metadata{Kind: "Secret", Name: v1secret.Name, Namespace: v1secret.Namespace,
		Labels: v1secret.Labels, Annotations: v1secret.Annotations}, Data: v1secret.Data}
	assert.Equal(t, &secret, NewSecret(v1secret))
}

func Test_NewSecretList(t *testing.T) {
	v1secretList := []*v1.Secret{getSecretV1()}
	secretList := []Secret{{Metadata: Metadata{Kind: "Secret", Name: v1secretList[0].Name,
		Namespace: v1secretList[0].Namespace, Labels: v1secretList[0].Labels, Annotations: v1secretList[0].Annotations},
		Data: v1secretList[0].Data}}
	assert.Equal(t, secretList, NewSecretList(v1secretList))
}

func Test_ToSecret_success(t *testing.T) {
	secret := Secret{Metadata: Metadata{Name: "test", Namespace: testNamespace,
		Annotations: map[string]string{"test": "test1"},
		Labels:      map[string]string{"test": "test2"}}, Data: map[string][]byte{"body": {16, 25, 68}}}
	v1secret := v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: secret.Name, Namespace: secret.Namespace,
		Labels: secret.Labels, Annotations: secret.Annotations}, Data: secret.Data}
	assert.Equal(t, &v1secret, secret.ToSecret())
}
