package entity

import (
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/extensions/v1beta1"
)

type ReplicaSet struct {
	Name            string
	CurrentVersion  string
	Replicas        int
	DesiredReplicas string
}

func NewReplicaSet(name string, currentVersion string, replicas int, desiredReplicas string) *ReplicaSet {
	return &ReplicaSet{name, currentVersion, replicas, desiredReplicas}
}

func TransformationReplicaSetListExtension(replicaSetList []v1beta1.ReplicaSet) *[]ReplicaSet {
	var transformedReplicaSetList []ReplicaSet
	var tmpReplicaSet *ReplicaSet
	for _, replicaSet := range replicaSetList {
		tmpReplicaSet = TransformationReplicaSetExtension(&replicaSet)
		transformedReplicaSetList = append(transformedReplicaSetList, *tmpReplicaSet)
	}
	return &transformedReplicaSetList
}

func TransformationReplicaSetListApp(replicaSetList []appsv1.ReplicaSet) *[]ReplicaSet {
	var transformedReplicaSetList []ReplicaSet
	var tmpReplicaSet *ReplicaSet
	for _, replicaSet := range replicaSetList {
		tmpReplicaSet = TransformationReplicaSetApp(&replicaSet)
		transformedReplicaSetList = append(transformedReplicaSetList, *tmpReplicaSet)
	}
	return &transformedReplicaSetList
}

func TransformationReplicaSetExtension(replicaSet *v1beta1.ReplicaSet) *ReplicaSet {
	return NewReplicaSet(replicaSet.Name, replicaSet.ObjectMeta.Annotations["deployment.kubernetes.io/revision"],
		int(replicaSet.Status.Replicas), replicaSet.Annotations["deployment.kubernetes.io/desired-replicas"])
}

func TransformationReplicaSetApp(replicaSet *appsv1.ReplicaSet) *ReplicaSet {
	return NewReplicaSet(replicaSet.Name, replicaSet.ObjectMeta.Annotations["deployment.kubernetes.io/revision"],
		int(replicaSet.Status.Replicas), replicaSet.Annotations["deployment.kubernetes.io/desired-replicas"])
}
