package main

import (
	"fmt"
	"time"
)

func main() {

	msgch := make(chan string, 10)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"

	close(msgch)

	// for {
	// 	msg, ok := <-msgch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("msg", msg)
	// }

	for msg := range msgch {
		fmt.Println(msg)
	}
	fmt.Println("Done")

}

func fetchRes(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("res %d", n)
}
