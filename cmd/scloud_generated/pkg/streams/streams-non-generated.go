// Package streams -- generated by scloudgen
// !! DO NOT EDIT !!
//
package streams

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/auth"
	model "github.com/splunk/splunk-cloud-sdk-go/services/streams"
)

// CreatePipeline
func CreatePipelineOverride(bypassValidation *bool, description *string, name string, filename string) (*model.PipelineResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.UplPipeline
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.CreatePipeline(model.PipelineRequest{Name: name, Data: data, BypassValidation: bypassValidation, Description: description})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ExpandPipelineOverride
func ExpandPipelineOverride(filename string) (*model.UplPipeline, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.UplPipeline
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.ExpandPipeline(data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PutGroupOverride
func PutGroupOverride(groupId string, filename string) (*model.GroupResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.GroupPutRequest
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.PutGroup(groupId, data)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

// PutTemplateOverride
func PutTemplateOverride(templateId string, filename string) (*model.TemplateResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.TemplatePutRequest
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.PutTemplate(templateId, data)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

// StartPreviewOverride
func StartPreviewOverride(filename string) (*model.PreviewStartResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.PreviewSessionStartRequest
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.StartPreview(data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateGroupOverride
func UpdateGroupOverride(groupId string, filename string) (*model.GroupResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.GroupPatchRequest
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.UpdateGroup(groupId, data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdatePipelineOverride
func UpdatePipelineOverride(id string, bypassValidation *bool, createUserId *string, description *string, name *string, filename string) (*model.PipelineResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	var data *model.UplPipeline
	if strings.Trim(filename, " ") != "" {
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bytes, &data)
		if err != nil {
			return nil, err
		}
	}

	resp, err := client.StreamsService.UpdatePipeline(id, model.PipelinePatchRequest{Data: data, Name: name, CreateUserId: createUserId, Description: description, BypassValidation: bypassValidation})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateTemplateOverride
func UpdateTemplateOverride(templateId string, description *string, name *string, filename string) (*model.TemplateResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	var data *model.UplPipeline
	if strings.Trim(filename, " ") != "" {
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bytes, &data)
		if err != nil {
			return nil, err
		}
	}

	resp, err := client.StreamsService.UpdateTemplate(templateId, model.TemplatePatchRequest{Data: data, Description: description, Name: name})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ValidatePipelineOverride
func ValidatePipelineOverride(filename string) (*model.ValidateResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.UplPipeline
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.ValidatePipeline(model.ValidateRequest{Upl: data})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CompileDSLOverride
func CompileDSLOverride(filename string) (*model.UplPipeline, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.DslCompilationRequest
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.CompileDSL(data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CompileSPLOverride
func CompileSPLOverride(filename string) (*model.UplPipeline, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.SplCompileRequest
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.CompileSPL(data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateTemplateOverride
func CreateTemplateOverride(description string, name string, filename string) (*model.TemplateResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.UplPipeline
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.CreateTemplate(model.TemplateRequest{Data: data, Name: name, Description: description})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MergePipelinesOverride
func MergePipelinesOverride(targetNode string, targetPort string, inputTreeFilename string, mainTreeFilename string) (*model.UplPipeline, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(inputTreeFilename)
	if err != nil {
		return nil, err
	}

	var inputtree model.UplPipeline
	err = json.Unmarshal(bytes, &inputtree)
	if err != nil {
		return nil, err
	}

	bytes, err = ioutil.ReadFile(mainTreeFilename)
	if err != nil {
		return nil, err
	}

	var maintree model.UplPipeline
	err = json.Unmarshal(bytes, &maintree)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.MergePipelines(model.PipelinesMergeRequest{InputTree: inputtree, MainTree: maintree, TargetNode: targetNode, TargetPort: targetPort})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PutConnectionOverride
func PutConnectionOverride(id string, filename string) (*model.ConnectionSaveResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.ConnectionPutRequest
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.PutConnection(id, data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateGroupOverride
func CreateGroupOverride(filename string) (*model.GroupResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.GroupRequest
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.CreateGroup(data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetOutputSchemaOverride
func GetOutputSchemaOverride(nodeUuid *string, sourcePortName *string, filename string) (map[string]model.UplType, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.UplPipeline
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.GetOutputSchema(model.GetOutputSchemaRequest{UplJson: data, NodeUuid: nodeUuid, SourcePortName: sourcePortName})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetInputSchemaOverride
func GetInputSchemaOverride(nodeUuid string, targetPortName string, filename string) (*model.UplType, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.UplPipeline
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.GetInputSchema(model.GetInputSchemaRequest{NodeUuid: nodeUuid, UplJson: data, TargetPortName: targetPortName})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateConnectionOverride
func UpdateConnectionOverride(id string, filename string) (*model.ConnectionSaveResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.ConnectionPatchRequest
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.UpdateConnection(id, data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
