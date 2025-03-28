// Package gototp is a Go package for generating Time-based One-Time Passwords (TOTP) conforming to RFC 6238.
// It supports multiple hash algorithms, including HMAC-SHA1, HMAC-SHA256, and HMAC-SHA512.
package gototp

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"hash"
	"math"
)

type algorithm int

const (
	SHA1 algorithm = iota
	SHA256
	SHA512
)

func (a algorithm) String() string {
	switch a {
	case SHA1:
		return "SHA1"
	case SHA256:
		return "SHA256"
	case SHA512:
		return "SHA512"
	default:
		return "Unknown Algorithm"
	}
}

type Totp struct {
	Secret    string
	Digits    int
	Period    int64
	Algorithm algorithm
	T0        int64
}

func (t Totp) GenerateTOTP(timestamp int64) (string, error) {
	// Calculate the time counter
	timeCounter := (timestamp - t.T0) / t.Period

	// Convert time counter to byte array (big endian)
	timeCounterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timeCounterBytes, uint64(timeCounter))

	// Generate hash
	hash, err := generateAlgorithm(t.Secret, timeCounterBytes, t.Algorithm)
	if err != nil {
		return "", err
	}

	// Extract dynamic binary code (truncate)
	offset := hash[len(hash)-1] & 0x0F
	binaryCode := int64(binary.BigEndian.Uint32(hash[offset:offset+4]) & 0x7FFFFFFF)

	// Calculate TOTP value
	totpValue := binaryCode % int64(math.Pow10(t.Digits))

	// Format TOTP value as a zero-padded string
	format := fmt.Sprintf("%%0%dd", t.Digits)
	return fmt.Sprintf(format, totpValue), nil
}

func generateAlgorithm(secret string, data []byte, algorithm algorithm) ([]byte, error) {
	var mac hash.Hash
	key := []byte(secret)

	switch algorithm {
	case SHA1:
		mac = hmac.New(sha1.New, key)
	case SHA256:
		mac = hmac.New(sha256.New, key)
	case SHA512:
		mac = hmac.New(sha512.New, key)
	default:
		return nil, fmt.Errorf("Unknown Algorithm")
	}

	mac.Write(data)
	return mac.Sum(nil), nil
}
