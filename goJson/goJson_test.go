package gojson

import (
	"testing"
)

func TestToJSON(t *testing.T) {
	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	test := []struct {
		name   string
		param  interface{}
		expect string
	}{
		{
			name:   "success - convert struct to json",
			param:  User{Name: "Alice", Age: 30, Email: "alice@example.com"},
			expect: `{"name":"Alice","age":30,"email":"alice@example.com"}`,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			goJSON := NewGoJSON()
			actual := goJSON.ToJSON(tt.param)
			if actual != tt.expect {
				t.Errorf("TestToJSON got = %v, want %v", actual, tt.expect)
			}
		})
	}
}

func TestToJSONPretty(t *testing.T) {
	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	test := []struct {
		name   string
		param  interface{}
		expect string
	}{
		{
			name:   "success - convert struct to json pretty",
			param:  User{Name: "Alice", Age: 30, Email: "alice@example.com"},
			expect: "{\n  \"name\": \"Alice\",\n  \"age\": 30,\n  \"email\": \"alice@example.com\"\n}",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			goJSON := NewGoJSON()
			actual := goJSON.ToJSONPretty(tt.param)
			if actual != tt.expect {
				t.Errorf("TestToJSON got = %v, want %v", actual, tt.expect)
			}
		})
	}
}

func TestToStruct(t *testing.T) {
	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	type params struct {
		jsonStr         string
		structInterface interface{}
	}

	user := new(User)

	test := []struct {
		name   string
		param  *params
		expect *User
		error  bool
	}{
		{
			name: "success - convert struct to struct",
			param: &params{
				jsonStr:         `{"name":"Alice","age":30,"email":"alice@example.com"}`,
				structInterface: user,
			},
			expect: user,
			error:  false,
		},
		{
			name: "success - convert struct to struct",
			param: &params{
				jsonStr:         `"name":"Alice","age":30,"email":"alice@example.com"}`,
				structInterface: user,
			},
			expect: user,
			error:  true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			goJSON := NewGoJSON()
			err := goJSON.ToStruct(tt.param.jsonStr, tt.param.structInterface)
			if err != nil != tt.error {
				t.Errorf("TestToStruct error: %v", err)
			}

			if tt.expect != tt.param.structInterface {
				t.Errorf("TestToStruct got = %v, want %v", tt.param.structInterface, tt.expect)
			}
		})
	}
}
