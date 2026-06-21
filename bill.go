package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   10.56,
	}

	return b
}

func (b bill) format() string {
	fs := "\n\t bill breakdown \n"
	total := 0.0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+": ", v)
		total += v
	}

	fs += fmt.Sprintf("%25v ...$%0.2f \n", "total: ", total)
	fs += fmt.Sprintf("%25v ...$%0.2f \n", "tip: ", b.tip)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addBill(name string, value float64) {
	b.items[name] = value
}

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("bill was saved!")
}
