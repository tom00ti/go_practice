package main

import "fmt"

func main() {
	var conferenceName = "Go Conffernces"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("ConferenceTeckes is %T, remaningTickets is %T, conferenceName is %T\n ", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tikets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get yout tickets here to attend")

	var userName string
	var userTikets int
	// ask user for their name

	userName = "Tom"
	userTikets = 2
	fmt.Printf("User %v booked  %v tickets.\n", userName, userTikets)

}
