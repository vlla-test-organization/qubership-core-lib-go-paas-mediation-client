package entity

import (
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func getServiceAccountV1() *v1.ServiceAccount {
	return &v1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: testNamespace,
		Annotations: map[string]string{"test": "test1"},
		Labels:      map[string]string{"test": "test2"}}, Secrets: []v1.ObjectReference{{Name: "s1"}, {Name: "s2"}}}
}

func Test_NewServiceAccount(t *testing.T) {
	serviceAccountV1 := getServiceAccountV1()
	serviceAccountTest := ServiceAccount{Metadata: Metadata{Kind: "ServiceAccount", Name: serviceAccountV1.Name,
		Namespace: serviceAccountV1.Namespace,
		Labels:    serviceAccountV1.Labels, Annotations: serviceAccountV1.Annotations}}
	for _, data := range serviceAccountV1.Secrets {
		secret := SecretInfo{Name: data.Name}
		serviceAccountTest.Secrets = append(serviceAccountTest.Secrets, secret)
	}
	assert.Equal(t, &serviceAccountTest, NewServiceAccount(serviceAccountV1))
}

func Test_NewServiceAccountList(t *testing.T) {
	serviceAccountV1List := []*v1.ServiceAccount{getServiceAccountV1()}
	serviceAccountTestList := []ServiceAccount{{Metadata: Metadata{Kind: "ServiceAccount",
		Name: serviceAccountV1List[0].Name, Namespace: serviceAccountV1List[0].Namespace,
		Labels: serviceAccountV1List[0].Labels, Annotations: serviceAccountV1List[0].Annotations},
		Secrets: []SecretInfo{{Name: "s1"}, {Name: "s2"}}}}
	assert.Equal(t, serviceAccountTestList, NewServiceAccountList(serviceAccountV1List))
}

func Test_ToServiceAccount(t *testing.T) {
	serviceAccountTest := ServiceAccount{Metadata: Metadata{Kind: "ServiceAccount", Name: "test",
		Namespace:   testNamespace,
		Annotations: map[string]string{"test": "test1"}, Labels: map[string]string{"test": "test2"}},
		Secrets: []SecretInfo{{Name: "s1"}, {Name: "s2"}}}
	assert.Equal(t, *getServiceAccountV1(), *serviceAccountTest.ToServiceAccount())
}
