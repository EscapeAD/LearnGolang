package main

import "fmt"

func main() {
	var large int
	var small int
	fmt.Println("Enter Large number first and followed by a small")
	fmt.Scan(&large, &small)
	// fmt.Println("Enter small number")
	// fmt.Scan(&small)
	fmt.Println("remainder is", large%small)
}
