package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTicket int = 50
	var remainingTicket uint = 50
	fmt.Printf("ConferenceTicket is %T. ramainingTicket is %T. conferenceName is %T\n",
		conferenceTicket, remainingTicket, conferenceName)

	var userName string
	var userTickets uint
	fmt.Println("Insert your name: ")
	fmt.Scan(&userName)
	fmt.Println("Insert number of ticket you want booking: ")
	fmt.Scan(&userTickets)
	remainingTicket = remainingTicket - userTickets
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)
}
