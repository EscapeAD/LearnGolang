package main

import (
	"fmt"
)

func main() {
	a := 22

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
}
