package main

import (
	"fmt"
	"github.com/michaelwp/lazygo/goBinaryConverter"
	"log"
)

func main() {
	num := int64(30)
	converter := goBinaryConverter.NewBinaryConverter()
	binary := converter.ToBinary(num)
	decimal, err := converter.ToDecimal(binary)
	if err != nil {
		log.Fatal("error converting binary to decimal:", err)
	}

	fmt.Printf("binaryConverter.ToBinary(%d) = %s\n", num, binary)
	fmt.Printf("binaryConverter.ToDecimal(%s) = %d\n", binary, decimal)
}
