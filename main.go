package main

import (
	"fmt"
	"math"
)

func greet(n string) {
	fmt.Println(n)
}

func bye(n string) {
	fmt.Println("good bye", n)
}

func multipleValue(n []string, f func(string)) {
	for _, val := range n {
		f(val)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}
func main() {

	greet("whatsupp")
	bye("mahir")

	multipleValue([]string{"aa", "bb", "cc"}, bye)
	fmt.Println(circleArea(10.9))

	sayHello("horous")

	for _, value := range points {
		fmt.Println(value)
	}

	mapTest()
}
