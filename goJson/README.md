# GoJson - Simple JSON Converter for Go

GoJson is a lightweight library designed to easily convert Go structs to JSON and vice versa. It provides simple methods to serialize and deserialize JSON data.

## 🚀 Features

- **Convert struct to JSON** (`ToJSON`)
- **Convert struct to pretty JSON** (`ToJSONIndent`)
- **Convert JSON to struct** (`ToStruct`)

## 📥 Installation

To install the package, run:

```sh
go get github.com/micahaelwp/goJson
```

## 📌 Usage

### Initialize the JSON Converter

```go
package main

import (
	"fmt"
	"github.com/michaelwp/goJson"
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
```

## 🏆 Why GoJsonStruct?

✅ Simple and easy-to-use API
✅ Supports pretty JSON formatting
✅ Minimal dependencies
✅ Improves Go JSON handling efficiency

## 📢 Contributing

Contributions are welcome! Feel free to submit pull requests or open issues.

---

Enjoy coding with **GoJsonStruct** – because working with JSON in Go should be effortless! 🚀
