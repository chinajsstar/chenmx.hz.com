package main

import "fmt"

func add(a int, b int) (int,int) {
	return a+b, a-b
}

func main() {
	sum, sub := add(1, 3)

	fmt.Println(sum, sub)
}


