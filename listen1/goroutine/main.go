package main

import (
	"fmt"
	"time"
)

func cale() {
	for i := 0; i < 10; i++ {
		time.Sleep(1*time.Second)
		fmt.Println("run ", i, " times")
	}

	fmt.Println("calc finish")
}

func main() {
	go cale()
	fmt.Println("i exited")

	time.Sleep(11*time.Second)
}
