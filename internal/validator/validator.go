//Filename internal/data/validator.go

package validator

import (
	"net/url"
	"regexp"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	PhoneRX = regexp.MustCompile(`^\+?\(?[0-9]{3}\)?\s?-\s?[0-9]{3}\s?-\s?[0-9]{4}$`)
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

// return true if a string match a specific regex pattern
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// check if string is a valid website
func ValidWebsite(website string) bool {
	_, err := url.ParseRequestURI(website)
	return err == nil
}

// check if errors map needs an entry
func (v *Validator) Check(ok bool, key string, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// check for repeating values in a slice
func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}
