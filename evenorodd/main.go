package main

import "fmt"

func main() {
	// define slice, loop through numbers 1 - 10, append them to slice
	intList := []int{}
	for i := 1; i <= 10; i++ {
		intList = append(intList, i)
	}

	// loop through slice, determining if even or odd and printing to user
	for _, num := range intList {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
