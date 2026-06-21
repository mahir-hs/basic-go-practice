package main

import "fmt"

var points = []int{1, 2, 3, 4, 5}

func sayHello(n string) {
	fmt.Println(n)
}

func mapTest() {
	menu := map[string]float64{
		"soup": 1.5,
		"pie":  5.9,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, " - ", v)
	}
}
