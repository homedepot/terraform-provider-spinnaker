package spinnaker

import (
   "encoding/json"
	"fmt"
	"regexp"
)

func validateApplicationName(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if !regexp.MustCompile(`^[a-zA-Z0-9-]+$`).MatchString(value) {
		errors = append(errors, fmt.Errorf("Only alphanumeric characters or '-' allowed in %q", k))
	}
	return
}

func validatePipelineJson(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
   if !json.Valid([]byte(value)) {
      errors = append(errors, fmt.Errorf("Not valid JSON"))
   }
	return
}

