package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make a new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"Pie": 45, "Pizza": 34},
		tip:   0,
	}
	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown\n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...%-25v\n", k+":", v)
		total += v
	}

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "total:", total)

	return fs
}

// update tip
func (b bill) updateTip(tip float64) {
	b.tip = tip
}

// add name
func (b bill) addItem(name string, price float64) {
	b.items[name] = price
}
