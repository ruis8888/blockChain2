package main

import (
	"fmt"
	"github.com/ruis8888/blockChain2/block"
)

func main() {

	b := block.NewBlock("", "gensis block")
	fmt.Println(b)
}
