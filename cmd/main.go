package main

import (
	"fmt"

	"github.com/uglo1/blockchain/utils"
)

func main() {
	fmt.Println(utils.IsFoundHost("127.0.0.1", 5001))
}
