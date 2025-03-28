// Package gojson is a lightweight library designed to easily convert Go structs to JSON and vice versa.
// It provides simple methods to serialize and deserialize JSON data.
package gojson

import (
	"encoding/json"
	"fmt"
)

// GoJSON interface
type GoJSON interface {
	ToJSON(v interface{}) string
	ToStruct(jsonStr string, v interface{}) error
	ToJSONPretty(v interface{}) string
}

// ImplGoJSONStruct struct
type ImplGoJSONStruct struct{}

// NewGoJSON create new GoJSON instance
func NewGoJSON() GoJSON {
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

// ToJSONPretty converts any struct to a prettified JSON string.
func (j ImplGoJSONStruct) ToJSONPretty(v interface{}) string {
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
