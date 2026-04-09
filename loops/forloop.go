package main

import "fmt"

func main() {
	sum := 0
	var i int
	for i = 1; i < 10; i++ {
		sum += i
		fmt.Printf("iteration no. %d - sum: %d\n", i, sum)
	}
	fmt.Printf("final sum: %d\n", sum)
}
