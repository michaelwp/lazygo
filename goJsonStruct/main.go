package goJsonStruct

import (
	"encoding/json"
	"fmt"
)

type GoJsonStruct interface {
	ToJSON(v interface{}) string
	ToStruct(jsonStr string, v interface{}) error
	ToJSONIndent(v interface{}) string
}

type ImplGoJsonStruct struct{}

func NewGoJsonStruct() GoJsonStruct {
	return ImplGoJsonStruct{}
}

// ToJSON converts any struct to a JSON string.
func (j ImplGoJsonStruct) ToJSON(v interface{}) string {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return "{}" // Return empty JSON if there's an error
	}
	return string(jsonData)
}

// ToJSONIndent converts any struct to a prettified JSON string.
func (j ImplGoJsonStruct) ToJSONIndent(v interface{}) string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(jsonData)
}

// FromJSON converts a JSON string back to a struct.
func (j ImplGoJsonStruct) ToStruct(jsonStr string, v interface{}) error {
	if err := json.Unmarshal([]byte(jsonStr), v); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}
	return nil
}
