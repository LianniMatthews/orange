//Filename internal/data/validator.go

package validator

import (
	"regexp"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// validation errors map

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

// Validator methods

// check if map has entries
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// add unique entry to map
func (v *Validator) AddError(key string, message string) {
	// check if key already exist
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// check if an element can be found in a list of items
func In(element string, list ...string) bool {
	for i := range list {
		if element == list[i] {
			return true
		}
	}
	return false
}
