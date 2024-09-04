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
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill
func format(b bill) string {
	fs := "Bill breakdown\n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprint("%-25v ...%-25v", k+":", v)
		total += v
	}

	//total
	fs += fmt.Sprintf("%-25 ...$%0.2f", "total:", total)

	return fs
}
