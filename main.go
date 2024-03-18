package main

import (
	"fmt"
	"fun-stuff/types"
	"fun-stuff/util"
)

// 1
func main() {
	user := types.User{
		Name: "John 2",
		Age:  18,
	}
	fmt.Printf("Hello, %+v \n", util.GetNumber(), user)
}
