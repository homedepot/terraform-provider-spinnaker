package api

import (
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"
	gate "github.com/spinnaker/spin/cmd/gateclient"
	gate_swagger "github.com/spinnaker/spin/gateapi"
)

func FormatAPIErrorMessage(function_name string, err error) error {
      switch err_sw := err.(type) {
      case gate_swagger.GenericSwaggerError:
         return fmt.Errorf("%s: %s() response payload: %s", err_sw.Error(), function_name, err_sw.Body())
      }
      return fmt.Errorf("%s: From %s()", err, function_name)
}

func CreatePipeline(client *gate.GatewayClient, pipeline interface{}) error {
	resp, err := client.PipelineControllerApi.SavePipelineUsingPOST(client.Context, pipeline, nil)

	if err != nil {
      return FormatAPIErrorMessage ("PipelineControllerApi.SavePipelineUsingPOST", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Encountered an error saving pipeline, status code: %d\n", resp.StatusCode)
	}

	return nil
}

func GetPipeline(client *gate.GatewayClient, applicationName, pipelineName string, dest interface{}) (map[string]interface{}, error) {
	jsonMap, resp, err := client.ApplicationControllerApi.GetPipelineConfigUsingGET(client.Context,
		applicationName,
		pipelineName)

	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return jsonMap, fmt.Errorf("%s", ErrCodeNoSuchEntityException)
		}
		return jsonMap, fmt.Errorf("Encountered an error getting pipeline %s, %s\n",
			pipelineName,
         FormatAPIErrorMessage ("PipelineControllerApi.GetPipelineConfigUsingGET", err))
	}

	if resp.StatusCode != http.StatusOK {
		return jsonMap, fmt.Errorf("Encountered an error getting pipeline in pipeline %s with name %s, status code: %d\n",
			applicationName,
			pipelineName,
			resp.StatusCode)
	}

	if jsonMap == nil {
		return jsonMap, fmt.Errorf(ErrCodeNoSuchEntityException)
	}

	if err := mapstructure.Decode(jsonMap, dest); err != nil {
		return jsonMap, err
	}

	return jsonMap, nil
}

func UpdatePipeline(client *gate.GatewayClient, pipelineID string, pipeline interface{}) error {
	_, resp, err := client.PipelineControllerApi.UpdatePipelineUsingPUT(client.Context, pipelineID, pipeline)

	if err != nil {
      return FormatAPIErrorMessage ("PipelineControllerApi.UpdatePipelineUsingPUT", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Encountered an error saving pipeline, status code: %d\n", resp.StatusCode)
	}

	return nil
}

func DeletePipeline(client *gate.GatewayClient, applicationName, pipelineName string) error {
	resp, err := client.PipelineControllerApi.DeletePipelineUsingDELETE(client.Context, applicationName, pipelineName)

	if err != nil {
      return FormatAPIErrorMessage ("PipelineControllerApi.DeletePipelineUsingDELETE", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Encountered an error deleting pipeline, status code: %d\n", resp.StatusCode)
	}

	return nil
}
