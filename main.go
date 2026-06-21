package main

import "fmt"

func updateName(x string) {
	x = "horus"
}

func updateNameByPointer(x *string) {
	*x = "mahir"
}

func updateMenu(y map[string]float64) {
	y["soup"] = 99.9
}
func main() {

	name := "titus"
	updateName(name)
	fmt.Println(name)

	menu := map[string]float64{
		"soup": 1.5,
		"pie":  5.9,
	}
	updateMenu(menu)
	fmt.Println(menu)

	m := &name

	fmt.Println(*m)

	updateNameByPointer(m)
	fmt.Println(name)

	myBill := newBill("mahir")
	fmt.Println(myBill.format())

	myBill.addBill("soup", 10.9)
	myBill.addBill("shorma", 80.5)

	myBill.updateTip(100)
	fmt.Println(myBill.format())
}
