package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conffernces"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings [50]string

	fmt.Printf("ConferenceTeckes is %T, remaningTickets is %T, conferenceName is %T\n ", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tikets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get yout tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTikets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter your number of tickets: ")
	fmt.Scan(&userTikets)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will confirmamtion email at %v \n", firstName, lastName, userTikets, email)

	remainingTickets = remainingTickets - userTikets
	bookings[0] = firstName + " " + lastName
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The array type: %T\n", bookings)
	fmt.Printf("The array length: %v\n", len(bookings))

	fmt.Printf("We have %v tickets remaining for %v\n", remainingTickets, conferenceName)
	//	fmt.Println(remainingTickets)
	//	fmt.Println(&remainingTickets)

	//userTikets = 2
	//fmt.Printf("User %v booked  %v tickets.\n", userName, userTikets)

}
