package main

import "fmt"

func main() {
	a := 31
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Print(b)
	fmt.Print(*b)

	*b = 21
	fmt.Println(a)
}
