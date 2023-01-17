//prints numbers from 1 to 100 that are divisible by 3
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print(i, ", ")
		}
	}
	fmt.Println()
}
