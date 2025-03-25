package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const METADATA_NAME = "metadataName1"

func getTestMetadata() Metadata {
	return Metadata{Name: METADATA_NAME, Annotations: map[string]string{"owner": "someOwner"}}
}

func getExpandTestMetadata() Metadata {
	testMetadata := getTestMetadata()
	testMetadata.Labels = map[string]string{}
	testMetadata.Generation = 1
	return testMetadata
}

func TestCastSecretToMetadata(t *testing.T) {
	testMetadata := getTestMetadata()
	secret := Secret{Metadata: testMetadata}
	assert.Equal(t, testMetadata, secret.GetMetadata())
}

func TestCastConfigMapToMetadata(t *testing.T) {
	testMetadata := getTestMetadata()
	configMap := ConfigMap{Metadata: testMetadata}
	assert.Equal(t, testMetadata, configMap.GetMetadata())
}

func TestCastRouteToMetadata(t *testing.T) {
	testMetadata := getTestMetadata()
	route := Route{Metadata: testMetadata}
	assert.Equal(t, testMetadata, route.GetMetadata())
}

func TestCastServiceAccountsToMetadata(t *testing.T) {
	testMetadata := getTestMetadata()
	serviceAccount := ServiceAccount{Metadata: testMetadata}
	assert.Equal(t, testMetadata, serviceAccount.GetMetadata())
}

func TestCastServicesToMetadata(t *testing.T) {
	testMetadata := getTestMetadata()
	service := Service{Metadata: testMetadata}
	assert.Equal(t, testMetadata, service.GetMetadata())
}

func TestCastPodsToMetadata(t *testing.T) {
	testMetadata := getTestMetadata()
	pod := Pod{Metadata: testMetadata}
	assert.Equal(t, testMetadata, pod.GetMetadata())
}

func TestCastNamespacesToMetadata(t *testing.T) {
	testMetadata := getTestMetadata()
	namespace := Namespace{Metadata: testMetadata}
	assert.Equal(t, testMetadata, namespace.GetMetadata())
}

func TestCastCertificatesToMetadata(t *testing.T) {
	testMetadata := getTestMetadata()
	certificate := Certificate{Metadata: testMetadata}
	assert.Equal(t, testMetadata, certificate.GetMetadata())
}

func TestNewMetadataFromInterface(t *testing.T) {
	testMetadata := getExpandTestMetadata()
	interfaceForMetadata := make(map[string]any)
	interfaceForMetadata["name"] = METADATA_NAME
	annotations := make(map[string]any)
	annotations["owner"] = "someOwner"
	interfaceForMetadata["annotations"] = annotations
	interfaceForMetadata["generation"] = int64(1)
	metadataFromInterface := NewMetadataFromInterface("", interfaceForMetadata)
	assert.Equal(t, testMetadata, metadataFromInterface)
}
