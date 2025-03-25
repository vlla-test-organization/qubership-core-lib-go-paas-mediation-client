package entity

import (
	v12 "github.com/openshift/api/apps/v1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
	"time"
)

func Test_NewDeployment_success(t *testing.T) {
	kuberDepl := v1.Deployment{
		Status: v1.DeploymentStatus{
			Conditions: []v1.DeploymentCondition{{}},
		},
	}
	kuberDeplExpected := &Deployment{
		Metadata: Metadata{Kind: "Deployment"},
		Status: DeploymentStatus{
			Conditions: []DeploymentCondition{{}},
		},
	}

	result := NewDeployment(&kuberDepl)

	assert.Equalf(t, kuberDeplExpected, result, "Not expected Deployment for income kuber deploymnet")
}

func Test_NewDeploymentConfig_success(t *testing.T) {
	openshiftDepl := v12.DeploymentConfig{
		Status: v12.DeploymentConfigStatus{
			Conditions: []v12.DeploymentCondition{{}},
		},
	}
	openshiftDeplExpected := &Deployment{
		Metadata: Metadata{Kind: "DeploymentConfig"},
		Spec:     DeploymentSpec{Replicas: &openshiftDepl.Spec.Replicas},
		Status: DeploymentStatus{
			Conditions: []DeploymentCondition{{}},
		},
	}
	result := NewDeploymentConfig(&openshiftDepl)

	assert.Equalf(t, openshiftDeplExpected, result, "Not expected Deployment for income openshift deploymnet")
}

func Test_givenNil_getFormattedTimeString_returnNil(t *testing.T) {
	var expected *string
	result := getFormattedTimeString(nil)
	assert.Equalf(t, expected, result, "Given nil time method should return nil")
}

func Test_givenZero_getFormattedTimeString_returnNil(t *testing.T) {
	var expected *string
	result := getFormattedTimeString(&metaV1.Time{})
	assert.Equalf(t, expected, result, "Given zero time method should return nil")
}

func Test_givenNow_getFormattedTimeString_success(t *testing.T) {
	now := metaV1.Now()
	result := getFormattedTimeString(&now)
	assert.Equalf(t, now.Format(time.RFC3339), *result, "Given now time method should return nil")
}
