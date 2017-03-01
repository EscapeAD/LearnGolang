package main

import "fmt"

func main() {
	a := "Hi I want to be store in memory"

	fmt.Println("a -", a)
	fmt.Println("a's memory address - ", &a)
}
