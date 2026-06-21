package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("choose option (a - add item, s - save bill, t- add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("item name: ", reader)
		price, _ := getInput("item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The Price must be a number!")
			promptOptions(b)
		}
		b.addBill(name, p)
		fmt.Println("item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("you choose tip: ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number!")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip updated - ", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the file - ", b.name)
	default:
		fmt.Println("you have entered wrong input, please try again!?")
		promptOptions(b)
	}
}

func main() {

	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)
}
