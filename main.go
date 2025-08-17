package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	var conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to ", conferenceName, "Booking application")
	fmt.Println("We have total of ", conferenceTickets, "tickets and ", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
	conferenceTickets = 30

}
