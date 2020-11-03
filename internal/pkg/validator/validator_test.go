//(c) Copyright 2020 Hewlett Packard Enterprise Development LP

package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twinj/uuid"

	gv "gopkg.in/validator.v2"
)

var (
	testUUID = uuid.NewV4().String()
)

func TestIDValidation(t *testing.T) {
	type idExample struct {
		ID string `json:"id" validate:"IDValidator"`
	}

	// A regular uuid should pass this test.
	s := idExample{ID: testUUID}
	err := Validate(s)
	assert.NoError(t, err)

	// Missing ID should fail.
	s = idExample{}
	err = Validate(s)
	assert.Error(t, err)
	assert.Equal(t, "ID: "+gv.ErrZeroValue.Error(), err.Error())

	// An empty ID should fail.
	s.ID = ""
	err = Validate(s)
	assert.Error(t, err)
	assert.Equal(t, "ID: "+gv.ErrZeroValue.Error(), err.Error())

	// Test for containsSpace error.
	test := []string{
		" " + testUUID + " abc",
		"\t" + testUUID,
		"\n" + testUUID,
		"\r" + testUUID,
	}
	for _, value := range test {
		s.ID = value
		err = Validate(s)
		assert.Error(t, err)
		assert.Equal(t, "ID: "+ErrSpaceInValue.Error(), err.Error())
	}
}
