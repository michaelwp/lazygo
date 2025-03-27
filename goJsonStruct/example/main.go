// gojsonstruct package example
package main

import (
	"fmt"
	gojsonstruct "github.com/michaelwp/lazygo/goJsonStruct"
)

func main() {
	jsonConverter := gojsonstruct.NewGoJSONStruct()

	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

	// Convert struct to JSON
	jsonStr := jsonConverter.ToJSON(user)
	fmt.Println(jsonStr) // Output: {"name":"Alice","age":30,"email":"alice@example.com"}

	// Convert struct to pretty JSON
	jsonPretty := jsonConverter.ToJSONIndent(user)
	fmt.Println(jsonPretty)

	// Convert json back to struct
	var userStruct User
	err := jsonConverter.ToStruct(jsonStr, &userStruct)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(userStruct) // Output: {Alice 30 alice@example.com}
	}
}
