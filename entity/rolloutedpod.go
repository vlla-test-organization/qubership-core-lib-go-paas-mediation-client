package entity

import (
	v1 "k8s.io/api/core/v1"
)

// todo re-design this!

type RolloutedPod struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// todo re-write in the next major release
func NewRolloutedPod(pod *v1.Pod) *RolloutedPod {
	var resultString string
	for _, condition := range pod.Status.Conditions {
		if condition.Type == "Ready" {
			if condition.Status == "True" {
				resultString = "Ready"
			} else {
				resultString = "Not ready"
			}
		}
	}
	return &RolloutedPod{pod.Name, resultString}
}

// todo re-write in the next major release
func NewRolledOutPod(pod Pod) *RolloutedPod {
	var resultString string
	for _, condition := range pod.Status.Conditions {
		if condition.Type == "Ready" {
			if condition.Status == "True" {
				resultString = "Ready"
			} else {
				resultString = "Not ready"
			}
		}
	}
	return &RolloutedPod{pod.Name, resultString}
}

func NewDeploymentsStatus(deploymentStatusMap map[string][]v1.Pod) *map[string][]RolloutedPod {
	mapForStruct := make(map[string][]RolloutedPod)
	for key, value := range deploymentStatusMap {
		var podList []RolloutedPod
		for i := range value {
			podList = append(podList, *NewRolloutedPod(&value[i]))
			if podList[i].Status != "Ready" {
				return nil
			}
		}
		mapForStruct[key] = podList
	}
	return &mapForStruct
}
