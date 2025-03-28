package main

import (
	"fmt"
	gojson "github.com/michaelwp/lazygo/goJson"
)

func main() {
	goJSON := gojson.NewGoJSON()

	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

	// Convert struct to JSON
	jsonStr := goJSON.ToJSON(user)
	fmt.Println(jsonStr) // Output: {"name":"Alice","age":30,"email":"alice@example.com"}

	// Convert struct to pretty JSON
	jsonPretty := goJSON.ToJSONPretty(user)
	fmt.Println(jsonPretty)

	// Convert json back to struct
	var userStruct User
	err := goJSON.ToStruct(jsonStr, &userStruct)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(userStruct) // Output: {Alice 30 alice@example.com}
	}
}
