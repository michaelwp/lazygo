# Binary Converter

A simple Go library for converting between decimal and binary numbers.

## Installation

```sh
go get github.com/michaelwp/binaryConverter
```

## Usage

Import the package and use the `BinaryConverter` interface to convert between decimal and binary formats.

### Example

```go
package main

import (
	"fmt"
	"github.com/michaelwp/binaryConverter"
)

func main() {
	converter := binaryConverter.NewBinaryConverter()

	// Convert decimal to binary
	binary := converter.ToBinary(42)
	fmt.Println("Binary of 42:", binary) // Output: "101010"

	// Convert binary to decimal
	decimal, err := converter.ToDecimal("101010")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Decimal of 101010:", decimal) // Output: 42
	}
}
```

## API Reference

### `ToBinary(decimal int64) string`
Converts a decimal number to a binary string.

- **Input:** `int64`
- **Output:** `string`

**Example:**
```go
converter.ToBinary(10) // Output: "1010"
```

### `ToDecimal(binary string) (int64, error)`
Converts a binary string to a decimal number.

- **Input:** `string` (must be a valid binary representation)
- **Output:** `int64, error`

**Example:**
```go
decimal, err := converter.ToDecimal("1010") // Output: 10, nil
```

## Error Handling
If the input binary string is invalid, `ToDecimal` returns an error.
```go
_, err := converter.ToDecimal("102")
if err != nil {
    fmt.Println("Invalid binary string:", err)
}
```

