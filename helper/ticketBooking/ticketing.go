package ticketBooking

import (
	"basics/helper"
	"fmt"
	"sync"
	"time"
)

var name string
var email string
var phone string
var ticketsPurchased uint
var remainingTickets uint = 50

// var bookings []string //this is a slice
// var bookings = [50]string{} //this is an array
// var bookings = make([]map[string]string, 0)
// var userData = make(map[string]string) //this is a map of string

var bookings = make([]helper.UserData, 0) // a slice of struct

// sync package provide asynchronouse functionality.
// waitGroup package waits for all goroutines/threads to finish before ending main thread
var waitGroup = sync.WaitGroup{}

func InitializeBooking() {

	fmt.Println("Hello welcome to my ticketing system.")

	for {
		fmt.Println("Enter your name")
		fmt.Scan(&name)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("Enter your phone number")
		fmt.Scan(&phone)

		fmt.Println("Enter the number of tickets you want to buy")
		fmt.Scan(&ticketsPurchased)

		var userData = helper.UserData{
			Name:                     name,
			PhoneNumber:              phone,
			Email:                    email,
			NumberOfTicketsPurchased: ticketsPurchased,
		}

		if soldOut() {
			fmt.Printf("There are only %v tickets available\n", remainingTickets)
			continue
		}

		remainingTickets = remainingTickets - ticketsPurchased

		bookings = append(bookings, userData)

		waitGroup.Add(1)                                        //add number of threads to wait for before ending main thread
		go sendEmailNotification(name, email, ticketsPurchased) //the go key word handles the sendEmailNotification method on a separate thread

		fmt.Printf("Thank you for buyinh %v tickets. You will receive a confirmation message. \n", ticketsPurchased)

		fmt.Printf("Remaining tickets: %v\n", remainingTickets)

		// for _, booking:=range bookings {

		// }

		fmt.Printf("Bookings made %v\n\n", bookings)

		if remainingTickets <= 0 {
			fmt.Println("There are no tickets available")
			break
		}

	}
	//WaitGroup method wait() is added to the lastline of the codes.
	//Blocks the main thread untilWaitGroup Caount is zero or unitl all threads are done executing.
	waitGroup.Wait()
}

func soldOut() bool {
	return ticketsPurchased > remainingTickets
}

func sendEmailNotification(name string, email string, ticketsPurchased uint) {
	time.Sleep(40 * time.Second) //simulating sending email.

	fmt.Println("\n#########")
	fmt.Printf("Sending email to %v. through %v\n", name, email)
	fmt.Printf("message: Hi %v, you have successfuly purcahsed %v tickets\n", name, ticketsPurchased)
	fmt.Println("########### \n")

	// done method is called to inicade this thread is done executing so the main thread doesnot have to wait any longer.
	waitGroup.Done()
}
