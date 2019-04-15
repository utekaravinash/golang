package main

import (
	"fmt"
)

// Add all numbers in the args
func Add(numbers ...int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func main() {
	sum := Add(1, 2, 3, 4)
	fmt.Println("Sum:", sum)
}
