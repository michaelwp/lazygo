# go_totp

`go_totp` is a Go package for generating Time-based One-Time Passwords (TOTP) conforming to RFC 6238. It supports multiple hash algorithms, including HMAC-SHA1, HMAC-SHA256, and HMAC-SHA512.

## Features

- Supports HMAC-SHA1, HMAC-SHA256, and HMAC-SHA512 algorithms.
- Configurable TOTP parameters such as digit length, time step, and initial counter value.
- Easy-to-use interface for generating TOTPs.

## Installation

To install the package, run:

```sh
go get github.com/michaelwp/go_totp
```

## Usage

Here's a step-by-step guide on how to use the `go_totp` library to generate a TOTP.

### Import the Package

First, import the package into your Go code:

```go
import "github.com/michaelwp/go-totp"
```

### Create a TOTP Instance

Create an instance of the `Totp` struct with the desired configuration. For example, to create a TOTP generator using HMAC-SHA256 with a 8-digit output:

```go
totp := go_totp.Totp{
    Secret:    "rahasia",
    Digits:    8,
    Period:    15,
    Algorithm: go_totp.SHA256,
    T0:        0,
}
```

### TOTP Struct

The `Totp` struct contains the following fields:

- `Secret`: The shared secret key used for generating the TOTP.
- `Digits`: The number of digits in the TOTP output.
- `Period`: The time step in seconds (default is 30 seconds).
- `Algorithm`: The hash algorithm to use (SHA1, SHA256, or SHA512).
- `T0`: The initial counter value (default is 0).

### Generating the TOTP

To generate the TOTP, call the `GenerateTOTP` method on the `Totp` instance, passing the current Unix timestamp as the argument. This method returns the generated TOTP as a string.

### Example

Here's a complete example demonstrating how to use the `go_totp` package:

```go
package main

import (
	"fmt"
	"github.com/michaelwp/go-totp"
	"log"
	"time"
)

func main() {
	// Get the current Unix time
	timestamp := time.Now().Unix()

	// define the variables
	totp := go_totp.Totp{
		Secret:    "rahasia",
		Digits:    8,
		Period:    15,
		Algorithm: go_totp.SHA256,
		T0:        0,
	}

	// Generate the TOTP
	pass, err := totp.GenerateTOTP(timestamp)
	if err != nil {
		log.Println("GenerateTOTP error:", err)
	}

	// Print the TOTP
	fmt.Println("TOTP:", pass)
}
```

### Running the Example

To run the example, save it to a file (e.g., `main.go`), and then run:

```sh
go run main.go
```

This will output a 8-digit TOTP based on the provided user ID and the current Unix time.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with your changes.

## Contact

For any questions or issues, please open an issue on GitHub.