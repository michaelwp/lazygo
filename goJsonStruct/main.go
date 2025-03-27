// Package gojsonstruct v1.1.0
package gojsonstruct

import (
	"encoding/json"
	"fmt"
)

// GoJSONStruct interface
type GoJSONStruct interface {
	ToJSON(v interface{}) string
	ToStruct(jsonStr string, v interface{}) error
	ToJSONIndent(v interface{}) string
}

// ImplGoJSONStruct is a receiver
type ImplGoJSONStruct struct{}

// NewGoJSONStruct create new Go JSON instance
func NewGoJSONStruct() GoJSONStruct {
	return ImplGoJSONStruct{}
}

// ToJSON converts any struct to a JSON string.
func (j ImplGoJSONStruct) ToJSON(v interface{}) string {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return "{}" // Return empty JSON if there's an error
	}
	return string(jsonData)
}

// ToJSONIndent converts any struct to a prettified JSON string.
func (j ImplGoJSONStruct) ToJSONIndent(v interface{}) string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(jsonData)
}

// ToStruct converts a JSON string back to a struct.
func (j ImplGoJSONStruct) ToStruct(jsonStr string, v interface{}) error {
	if err := json.Unmarshal([]byte(jsonStr), v); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}
	return nil
}
