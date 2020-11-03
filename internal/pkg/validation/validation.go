//(c) Copyright 2020 Hewlett Packard Enterprise Development LP

// validator provides utilities used for data validator.

package validation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ValidateStructFields tests a generic struct for the existence
// of required fields. For string fields it also returns an error
// if a required field is empty.
// The use case for this helper function is in validating input from
// handlers that have many fields.
// Note that for structs that have json tags defined, the requiredFieldIDs
// need to be the json field names.
func ValidateStructFields(in interface{}, requiredFieldIDs []string) (err error) {
	var inAsMap map[string]interface{}
	temp, err := json.Marshal(in)
	if err != nil {
		return errors.New("error validating input struct")
	}
	err = json.Unmarshal(temp, &inAsMap)
	if err != nil {
		return errors.New("error validating input struct")
	}

	for _, requiredFieldID := range requiredFieldIDs {
		// Make sure the field is in the data.
		if val, ok := inAsMap[requiredFieldID]; !ok || len(fmt.Sprintf("%v", val)) == 0 {
			return errors.New("required input field " + requiredFieldID + " not specified")
		}
	}

	return nil
}
