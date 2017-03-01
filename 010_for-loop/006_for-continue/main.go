package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			// Restart loop or continue with new loop from this point
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}

}
