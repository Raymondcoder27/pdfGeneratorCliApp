package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(n string) bill {
	b := bill{
		name:  "",
		items: map[string]float64{},
		tip:   0,
	}

}
