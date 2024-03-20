package main

import (
	"fmt"
	"strings"
)

type TransFormFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s) + "0000"
}

func Prefixer(prefix string) TransFormFunc {
	return func(s string) string {
		return prefix + "-" + s
	}

}

func transformString(s string, fn TransFormFunc) string {
	return fn(s)
}
func main() {
	fmt.Println(transformString("hello Sailor!", Uppercase))
	fmt.Println(transformString("hello Sailor!", Prefixer("PFoo")))

}
