package main

import (
	"fmt"

	"github.com/whosydd/go_utils"
)

func main() {
	dst := "../LICENSE.md"
	flag, err := go_utils.IsExist(dst)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("flag: %v\n", flag)
}
