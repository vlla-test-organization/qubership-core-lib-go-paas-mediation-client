package entity

import (
	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func Test_TransformationReplicaSetExtension(t *testing.T) {
	replica := &v1beta1.ReplicaSet{ObjectMeta: v1.ObjectMeta{Name: testReplicaSet,
		Annotations: map[string]string{"deployment.kubernetes.io/revision": "1",
			"deployment.kubernetes.io/desired-replicas": "3"}},
		Status: v1beta1.ReplicaSetStatus{Replicas: 2}}
	newReplica := &ReplicaSet{Name: replica.Name, CurrentVersion: "1", Replicas: 2, DesiredReplicas: "3"}
	assert.Equal(t, newReplica, TransformationReplicaSetExtension(replica))
}

func Test_TransformationReplicaSetListExtension(t *testing.T) {
	replica := v1beta1.ReplicaSet{ObjectMeta: v1.ObjectMeta{Name: testReplicaSet,
		Annotations: map[string]string{"deployment.kubernetes.io/revision": "1",
			"deployment.kubernetes.io/desired-replicas": "3"}},
		Status: v1beta1.ReplicaSetStatus{Replicas: 2}}
	replicaList := []v1beta1.ReplicaSet{replica}
	newReplicaList := &[]ReplicaSet{{Name: replica.Name, CurrentVersion: "1", Replicas: 2, DesiredReplicas: "3"}}
	assert.Equal(t, newReplicaList, TransformationReplicaSetListExtension(replicaList))
}

func Test_TransformationReplicaSetApp(t *testing.T) {
	replica := &appsv1.ReplicaSet{ObjectMeta: v1.ObjectMeta{Name: testReplicaSet,
		Annotations: map[string]string{"deployment.kubernetes.io/revision": "1",
			"deployment.kubernetes.io/desired-replicas": "3"}},
		Status: appsv1.ReplicaSetStatus{Replicas: 2}}
	newReplica := &ReplicaSet{Name: replica.Name, CurrentVersion: "1", Replicas: 2, DesiredReplicas: "3"}
	assert.Equal(t, newReplica, TransformationReplicaSetApp(replica))
}

func Test_TransformationReplicaSetListApp(t *testing.T) {
	replica := appsv1.ReplicaSet{ObjectMeta: v1.ObjectMeta{Name: testReplicaSet,
		Annotations: map[string]string{"deployment.kubernetes.io/revision": "1",
			"deployment.kubernetes.io/desired-replicas": "3"}},
		Status: appsv1.ReplicaSetStatus{Replicas: 2}}
	replicaList := []appsv1.ReplicaSet{replica}
	newReplicaList := &[]ReplicaSet{{Name: replica.Name, CurrentVersion: "1", Replicas: 2, DesiredReplicas: "3"}}
	assert.Equal(t, newReplicaList, TransformationReplicaSetListApp(replicaList))
}
