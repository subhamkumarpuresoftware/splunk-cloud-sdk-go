// Package streams -- generated by scloudgen
// !! DO NOT EDIT !!
//
package streams

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	model "github.com/splunk/splunk-cloud-sdk-go/services/streams"
)

type HttpResponse struct {
	Code    *string                `json:"code,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Message *string                `json:"message,omitempty"`
}

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

	var data model.Pipeline
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
func StartPreviewOverride(recordsLimit *int32, recordsPerPipeline *int32, sessionLifetimeMsfilename *int64, filename string) (*model.PreviewStartResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data model.Pipeline
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resp, err := client.StreamsService.StartPreview(model.PreviewSessionStartRequest{RecordsLimit: recordsLimit, RecordsPerPipeline: recordsPerPipeline, SessionLifetimeMs: sessionLifetimeMsfilename, Upl: data})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdatePipelineOverride
func UpdatePipelineOverride(id string, bypassValidation *bool, description *string, name string, filename string) (*model.PipelineResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	var data model.Pipeline
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

	resp, err := client.StreamsService.UpdatePipeline(id, model.PipelineRequest{Data: data, Name: name, Description: description, BypassValidation: bypassValidation})
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

	var data *model.Pipeline
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

	var data model.Pipeline
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

// CompileOverride
func CompileOverride(validate *bool, filename string) (*model.Pipeline, error) {
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

	data.Validate = validate

	resp, err := client.StreamsService.Compile(data)
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

	var data model.Pipeline
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

	var data model.Pipeline
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

	var data model.Pipeline
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

func PatchPipelineOverride(id string, bypassValidation *bool, createUserId *string, description *string, name *string, filename string) (*model.PipelineResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}
	var data *model.Pipeline
	if filename != "" {

		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		//var data *model.Pipeline
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			return nil, err
		}
	}

	resp, err := client.StreamsService.PatchPipeline(id, model.PipelinePatchRequest{Data: data, BypassValidation: bypassValidation, CreateUserId: createUserId, Description: description, Name: name})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func UploadFileOverride(arg string) (*HttpResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	fileName := arg

	if _, err := os.Stat(fileName); err != nil {
		return nil, err
	}

	err = client.StreamsService.UploadFile(fileName)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
