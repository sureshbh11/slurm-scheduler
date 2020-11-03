//(c) Copyright 2020 Hewlett Packard Enterprise Development LP

package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateStructFields(t *testing.T) {
	type ToTest struct {
		StrID    string
		IntValue int
	}

	input := ToTest{"foo", 42}
	err := ValidateStructFields(&input, []string{"StrID", "IntValue"})
	assert.Nil(t, err)

	input = ToTest{StrID: "foo"}
	err = ValidateStructFields(&input, []string{"StrID", "IntValue"})
	assert.Nil(t, err)

	input = ToTest{StrID: ""}
	err = ValidateStructFields(&input, []string{"StrID", "IntValue"})
	assert.NotNil(t, err)
}

func TestValidateStructFieldsJSON(t *testing.T) {
	type ToTest struct {
		StrID    string `json:"str_id"`
		IntValue int    `json:"int_value"`
	}

	input := ToTest{"foo", 42}
	err := ValidateStructFields(&input, []string{"str_id", "int_value"})
	assert.Nil(t, err)

	input = ToTest{StrID: "foo"}
	err = ValidateStructFields(&input, []string{"str_id", "int_value"})
	assert.Nil(t, err)

	input = ToTest{StrID: ""}
	err = ValidateStructFields(&input, []string{"str_id", "int_value"})
	assert.NotNil(t, err)
}
