package main

import (
	"fmt"
	"fun-stuff/types"
	"fun-stuff/util"
)

// 1
func main() {
	user := types.User{
		Name: "John",
		Age:  18,
	}
	fmt.Printf("Hello, %+v \n", util.GetNumber(), user)
}
