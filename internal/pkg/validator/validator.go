//(c) Copyright 2020 Hewlett Packard Enterprise Development LP

package validator

import (
	"errors"
	"reflect"
	"regexp"

	gv "gopkg.in/validator.v2"
)

var (
	// ErrSpaceInValue is an error returned when a value contains a space.
	ErrSpaceInValue = gv.TextErr{Err: errors.New("space in value")}
)

// Validate sets up the input validator methods and then attempts to validate the input
// parameter. For optimization we could make a generator and initialize it
// only once within the http-client.go file and then make it a component of the views.
func Validate(input interface{}) error {
	v := gv.NewValidator()

	_ = v.SetValidationFunc("IDValidator", idValidator)

	return v.Validate(input)
}

func idValidator(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return gv.ErrUnsupported
	}

	value := st.String()

	if err := isNotZeroStringValidator(value); err != nil {
		return err
	}

	if err := containsNoSpaceValidator(value); err != nil {
		return err
	}

	return nil
}

// containsNoSpaceValidator checks if the string contains a whitespace
// character. If so, an error is returned.
func containsNoSpaceValidator(value string) error {
	space := regexp.MustCompile(`\s+`)
	cleaned := space.ReplaceAllString(value, "")
	if len(cleaned) != len(value) {
		return ErrSpaceInValue
	}
	return nil
}

// isNotZeroStringValidator checks if the string has at least one character.
// This could be accomplished using the nonzero tag.
func isNotZeroStringValidator(value string) error {
	if len(value) == 0 {
		return gv.ErrZeroValue
	}
	return nil
}
