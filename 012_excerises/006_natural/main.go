package main

import "fmt"

func main() {
	number := 0
	for i := 0; i <= 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			number += i
		}
	}
	fmt.Println(number)
}
