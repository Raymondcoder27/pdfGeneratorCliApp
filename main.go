package main

import "fmt"

func main() {
	myBill := newBill("My Bill")
	myBill.updateTip(40)
	myBill.addItem("onions", 23)
	myBill.addItem("carrot", 49)
	myBill.addItem("pineapple", 14)

	fmt.Println(myBill.format())
}
