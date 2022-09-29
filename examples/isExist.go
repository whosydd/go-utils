package main

import (
	"fmt"

	"github.com/whosydd/goutils"
)

func main() {
	dst := "../LICENSE.md"
	flag, err := goutils.IsExist(dst)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("flag: %v\n", flag)
}
