package entity

import (
	"github.com/netcracker/qubership-core-lib-go-paas-mediation-client/v8/types"
)

// todo in next major release change to
//type DeploymentResponse struct {
//	ReplicaSet            []DeploymentRollout `json:"replicaSet"`
//	ReplicationController []DeploymentRollout `json:"replicationController"`
//}

type DeploymentResponse struct {
	Deployments []DeploymentRollout `json:"deployments"`
}

type DeploymentRollout struct {
	Name    string `json:"name"`
	Kind    string `json:"kind"`
	Active  string `json:"active"`
	Rolling string `json:"rolling"`
}

func NewDeploymentResponse(deployments []*DeploymentRollout) *DeploymentResponse {
	deploymentsSlice := make([]DeploymentRollout, 0)
	for _, d := range deployments {
		if d != nil {
			deploymentsSlice = append(deploymentsSlice, *d)
		}
	}
	return &DeploymentResponse{Deployments: deploymentsSlice}
}

func NewDeploymentRolloutResponseObj(deploymentName string, rollingName string, activeName string) *DeploymentRollout {
	return &DeploymentRollout{Name: deploymentName, Kind: types.ReplicaSet, Active: activeName, Rolling: rollingName}
}

func NewDeploymentConfigRolloutResponseObj(deploymentConfigName string, rollingName string, activeName string) *DeploymentRollout {
	return &DeploymentRollout{Name: deploymentConfigName, Kind: types.ReplicationController, Active: activeName, Rolling: rollingName}
}
