package exportentities

import (
	"github.com/ovh/cds/sdk"
)

// WorkerModel is the as code format of a worker model
type WorkerModel struct {
	Name          string            `json:"name" yaml:"name"`
	Group         string            `json:"group" yaml:"group"`
	Communication string            `json:"communication,omitempty" yaml:"communication,omitempty"`
	Provision     int64             `json:"provision,omitempty" yaml:"provision,omitempty"`
	Image         string            `json:"image" yaml:"image"`
	Description   string            `json:"description" yaml:"description"`
	Type          string            `json:"type" yaml:"type"`
	Flavor        string            `json:"flavor,omitempty" yaml:"flavor,omitempty"`
	Envs          map[string]string `json:"envs,omitempty" yaml:"envs,omitempty"`
	PatternName   string            `json:"pattern_name,omitempty" yaml:"pattern_name,omitempty"`
	Shell         string            `json:"shell,omitempty" yaml:"shell,omitempty"`
	PreCmd        string            `json:"pre_cmd,omitempty" yaml:"pre_cmd,omitempty"`
	Cmd           string            `json:"cmd,omitempty" yaml:"cmd,omitempty"`
	PostCmd       string            `json:"post_cmd,omitempty" yaml:"post_cmd,omitempty"`
	Restricted    bool              `json:"restricted,omitempty" yaml:"restricted,omitempty"`
	IsDeprecated  bool              `json:"is_deprecated,omitempty" yaml:"is_deprecated,omitempty"`
}

// NewWorkerModel creates an exportentities WorkerModel from a struct sdk.Model
func NewWorkerModel(wm sdk.Model) WorkerModel {
	model := WorkerModel{
		Type:          wm.Type,
		Name:          wm.Name,
		PatternName:   wm.PatternName,
		Communication: wm.Communication,
		Group:         wm.Group.Name,
		IsDeprecated:  wm.IsDeprecated,
		Provision:     wm.Provision,
		Description:   wm.Description,
		Restricted:    wm.Restricted,
		Image:         wm.Image,
	}

	switch wm.Type {
	case sdk.Docker:
		model.Shell = wm.ModelDocker.Shell
		model.Image = wm.ModelDocker.Image
		model.Cmd = wm.ModelDocker.Cmd
		model.Envs = wm.ModelDocker.Envs
	case sdk.VSphere, sdk.Openstack:
		model.Flavor = wm.ModelVirtualMachine.Flavor
		model.Image = wm.ModelVirtualMachine.Image
		model.PreCmd = wm.ModelVirtualMachine.PreCmd
		model.Cmd = wm.ModelVirtualMachine.Cmd
		model.PostCmd = wm.ModelVirtualMachine.PostCmd
	}

	return model
}

// GetWorkerModel convert an exportentities to a real sdk.Model
func (wm WorkerModel) GetWorkerModel() (sdk.Model, error) {
	if err := wm.IsValid(); err != nil {
		return sdk.Model{}, err
	}

	model := sdk.Model{
		Type:          wm.Type,
		Name:          wm.Name,
		PatternName:   wm.PatternName,
		Communication: wm.Communication,
		Group:         sdk.Group{Name: wm.Group},
		IsDeprecated:  wm.IsDeprecated,
		Provision:     wm.Provision,
		Description:   wm.Description,
		Restricted:    wm.Restricted,
	}

	switch wm.Type {
	case sdk.Docker:
		model.ModelDocker = sdk.ModelDocker{
			Shell: wm.Shell,
			Image: wm.Image,
			Cmd:   wm.Cmd,
			Envs:  wm.Envs,
		}
	case sdk.VSphere, sdk.Openstack:
		model.ModelVirtualMachine = sdk.ModelVirtualMachine{
			Image:   wm.Image,
			Flavor:  wm.Flavor,
			Cmd:     wm.Cmd,
			PostCmd: wm.PostCmd,
			PreCmd:  wm.PreCmd,
		}
	}

	return model, nil
}

// IsValid returns error if worker model is invalid
func (wm *WorkerModel) IsValid() error {
	if wm.Name == "" {
		return sdk.NewErrorFrom(sdk.ErrWrongRequest, "Error: worker model name is not provided")
	}

	if wm.Group == "" {
		return sdk.NewErrorFrom(sdk.ErrWrongRequest, "Error: group is not provided")
	}

	switch wm.Type {
	case sdk.Docker:
		if wm.Image == "" {
			return sdk.NewErrorFrom(sdk.ErrWrongRequest, "Error: Docker image not provided")
		}
		if wm.PatternName == "" {
			if wm.Shell == "" {
				return sdk.NewErrorFrom(sdk.ErrWrongRequest, "Error: main shell command not provided")
			}
			if wm.Cmd == "" {
				return sdk.NewErrorFrom(sdk.ErrWrongRequest, "Error: main worker command not provided")
			}
		}
	case sdk.Openstack:
		if wm.Image == "" {
			return sdk.NewErrorFrom(sdk.ErrWrongRequest, "Error: Openstack image not provided")
		}
		if wm.Flavor == "" {
			return sdk.NewErrorFrom(sdk.ErrWrongRequest, "Error: Openstack flavor not provided")
		}
		if wm.PatternName == "" {
			if wm.Cmd == "" {
				return sdk.NewErrorFrom(sdk.ErrWrongRequest, "Error: Openstack command not provided")
			}
		}
	case sdk.VSphere:
		if wm.Image == "" {
			return sdk.NewErrorFrom(sdk.ErrWrongRequest, "Error: VSphere image not provided")
		}
		if wm.PatternName == "" {
			if wm.Cmd == "" {
				return sdk.NewErrorFrom(sdk.ErrWrongRequest, "Error: VSphere main worker command empty")
			}
		}
	default:
		return sdk.NewErrorFrom(sdk.ErrWrongRequest, "Unknown worker type: %s", wm.Type)
	}

	return nil
}
