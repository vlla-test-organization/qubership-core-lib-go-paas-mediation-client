package entity

const (
	AppNameProp          string = "app_name"
	AppVersionProp       string = "app_version"
	NameProp             string = "name"
	FamilyNameProp       string = "family_name"
	VersionProp          string = "version"
	BlueGreenVersionProp string = "bluegreen_version"
	StateProp            string = "state"
)

type DeploymentFamilyVersion struct {
	AppName          string `json:"app_name"`
	AppVersion       string `json:"app_version"`
	Name             string `json:"name"`
	FamilyName       string `json:"family_name"`
	BlueGreenVersion string `json:"bluegreen_version"`
	Version          string `json:"version"`
	State            string `json:"state"`
}

func DeploymentToDeploymentFamilyVersion(labels map[string]string) DeploymentFamilyVersion {
	return DeploymentFamilyVersion{
		AppName:          labels[AppNameProp],
		AppVersion:       labels[AppVersionProp],
		Name:             labels[NameProp],
		FamilyName:       labels[FamilyNameProp],
		BlueGreenVersion: labels[BlueGreenVersionProp],
		Version:          labels[VersionProp],
		State:            labels[StateProp],
	}
}
