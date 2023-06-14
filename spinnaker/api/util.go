package api

import (
	"fmt"
	gate_swagger "github.com/spinnaker/spin/gateapi"
)

func FormatAPIErrorMessage(function_name string, err error) error {
      switch err_sw := err.(type) {
      case gate_swagger.GenericSwaggerError:
         return fmt.Errorf("%s: %s() response payload: %s", err_sw.Error(), function_name, err_sw.Body())
      }
      return fmt.Errorf("%s: From %s()", err, function_name)
}

