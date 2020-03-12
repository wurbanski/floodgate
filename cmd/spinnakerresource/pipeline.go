package spinnakerresource

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codilime/floodgate/cmd/gateclient"
	"github.com/codilime/floodgate/cmd/util"
)

// Pipeline object
type Pipeline struct {
	*Resource
	name        string
	application string
}

// Init initialize pipeline
func (p *Pipeline) Init(api *gateclient.GateapiClient, localData map[string]interface{}) error {
	if err := p.validate(localData); err != nil {
		return err
	}
	name := localData["name"].(string)
	application := localData["application"].(string)
	localState, err := json.Marshal(localData)
	if err != nil {
		return err
	}
	p.Resource = &Resource{
		localState:   localState,
		spinnakerAPI: api,
	}
	p.name = name
	p.application = application
	if err := p.loadRemoteState(); err != nil {
		return err
	}
	return nil
}

// LoadRemoteState get remote resource
func (p *Pipeline) loadRemoteState() error {
	successPayload, resp, err := p.spinnakerAPI.ApplicationControllerApi.GetPipelineConfigUsingGET(p.spinnakerAPI.Context, p.application, p.name)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Encountered an error getting pipeline in pipeline %s with name %s, status code: %d", p.application, p.name, resp.StatusCode)
	}

	jsonPayload, err := json.Marshal(successPayload)
	if err != nil {
		return err
	}
	p.remoteState = jsonPayload

	return nil
}

// SaveRemoteState is used to save state remotely
func (p Pipeline) SaveRemoteState() error {
	var jsonPayload interface{}
	err := json.Unmarshal(p.localState, &jsonPayload)
	if err != nil {
		return err
	}

	saveResp, err := p.spinnakerAPI.PipelineControllerApi.SavePipelineUsingPOST(p.spinnakerAPI.Context, jsonPayload)
	if err != nil {
		return err
	}
	if saveResp.StatusCode != http.StatusOK {
		return fmt.Errorf("Encountered an error saving pipeline, status code: %d", saveResp.StatusCode)
	}

	return nil
}

func (p Pipeline) validate(localData map[string]interface{}) error {
	if err := util.AssertMapKeyIsString(localData, "name", true); err != nil {
		return err
	}
	if err := util.AssertMapKeyIsString(localData, "application", true); err != nil {
		return err
	}
	// template is optional key, but it requires defined schema
	err := util.AssertMapKeyIsStringMap(localData, "template", true)
	if err == nil {
		template := localData["template"].(map[string]interface{})
		if err := util.AssertMapKeyIsString(template, "schema", true); err != nil {
			return err
		}
	}
	return nil
}
