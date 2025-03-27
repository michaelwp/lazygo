package goBinaryConverter

import (
	"fmt"
	"math/big"
	"strconv"
)

type BinaryConverter interface {
	ToBinary(decimal int64) string
	ToDecimal(binary string) (int64, error)
}

type Impl struct{}

func NewBinaryConverter() BinaryConverter {
	return Impl{}
}

func (i Impl) ToBinary(decimal int64) string {
	return strconv.FormatInt(decimal, 2)
}

func (i Impl) ToDecimal(binary string) (int64, error) {
	decimal, isTrue := new(big.Int).SetString(binary, 2)
	if !isTrue {
		return 0, fmt.Errorf("binaryConverter.ToDecimal(%s) is invalid", binary)
	}

	return decimal.Int64(), nil
}
