package main

import "fmt"

func main() {
	a := "save this in memory"

	fmt.Println(a)
	fmt.Println(&a) //save to memory address

	var b *string = &a

	fmt.Println(b)
	fmt.Println(*b) //* as a operator
}
