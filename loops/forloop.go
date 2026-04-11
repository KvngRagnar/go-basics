package main

import "fmt"

func main() {
	count := 5
	fmt.Println("countdown")

	for count > 0 {
		fmt.Println(count)
		count--
	}
	fmt.Println("Kaboooom")
}
