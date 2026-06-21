package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	names := []string{"aa", "bb", "cc"}

	for index, value := range names {
		fmt.Println(index, value)
	}

	if x <= 0 {
		fmt.Println("value is not changed!")
	} else if x > 0 {
		fmt.Println("value got changed")
	}
}
