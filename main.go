package main

import (
	"basics/helper"
	"basics/helper/ticketBooking"
	"fmt"
)

// struct
type Human struct {
	name string
	age  int32
	job  string
}

// use := for variable declarations and should be used inside functions

func main1() {

	// var age, heigh int = 10, 15

	// var sliceOfMaps = make([]map[string]uint32, 0)

	var age uint = 10
	var height float32 = 15.5
	var userName string

	fmt.Println(helper.AppName)

	fmt.Println("Enter your name")

	fmt.Scan(&userName) //& is a special character called pointer. it access the memory loccation of the variable

	fmt.Println(mySelf(userName, age, height))
	my(10.5)

	var person Human
	person.name = "Kweku"

	// mapBasics()

}

// this function returns two values
func mySelf(userName string, age uint, height float32) (me string, food string) {
	me = fmt.Sprintf("My name is %v. I'm %v years old. I am %v feet tall.", userName, age, height)
	food = "I like beans."
	return
}

func my(amount float32) {
	// Array
	// var ages = [...]int{15, 20, 13}
	var names = [3]string{"Adwoa", "Araba", "Philip"}

	// Slices
	var expenses = []float32{15.20, 10.50, 85.00}
	// var profit = []float32{18.00, 25.50, 100.00}

	// var total float32 = append(expenses, profit...)

	// switch statement
	switch amount {
	case 10.50, 50.00:
		fmt.Println("High Amount")
	case 18.00, 25.00:
		fmt.Println("Low Amounts")
	default:
		fmt.Println("Uncategorized amount")
	}

	for i := 0; i < len(expenses); i++ {
		fmt.Println("Expense: ", expenses[i])
	}

	// for loop using range
	for index, value := range names {
		fmt.Println(index, value)
	}
}

func main() {
	ticketBooking.InitializeBooking()
}
