package main

import "fmt"

func main() {

	// var ages [3]int = [3]int{10, 17, 26}
	var ages = [3]int{10, 17, 26}

	names := [4]string{"aa", "bb", "cc", "dd"}
	names[1] = "gg"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	scores := []int{100, 200, 300}
	scores = append(scores, 400)
	fmt.Println(scores, len(scores))

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, len(rangeOne), rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "kk")
	fmt.Println(rangeOne)
}
