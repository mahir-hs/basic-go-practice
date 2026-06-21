package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	greeting := "hello world"
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hola!"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	fmt.Println("Original string = ", greeting)

	ages := []int{4, 10, 5, 9, 55, 78, 6, 7, 7}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 7)
	fmt.Println(index)
}
