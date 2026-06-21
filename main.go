package main

import "fmt"

func main() {

	//strings
	var nameOne string = "i am lost,"
	var nameTwo = "so lost..."
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "oh no!!"
	nameThree = "i am lost"

	fmt.Println(nameOne, nameTwo, nameThree)

	age := 50
	name := "titus"

	fmt.Println("my name is ", name, " and my age is ", age)
	fmt.Println("my name is _, my age is _", name, age)
	fmt.Printf("my name is %v and my age is %v \n", name, age)
	fmt.Printf("my name is %q and my age is %q \n", name, age)
	fmt.Printf("age is type of %T \n", age)
	fmt.Printf("you scorred %0.2f marks \n", 500.159)

	str := fmt.Sprintf("my name is %v and my age is %v \n", name, age)

	fmt.Println(str)
}
