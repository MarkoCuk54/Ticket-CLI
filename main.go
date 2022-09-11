package main

import "fmt"

func main() {
	const conferenceName = "Go Conference"
	const tickets = 50
	ticketsLeft := 50

	fmt.Printf("Welcome to the %v.\n", conferenceName)
	fmt.Printf("You can get youre Tickets here! There are %v tickets and %v are left.\n", tickets, ticketsLeft)

	bookings := []string{"HQ"}
	var userName string
	var userTickets int8

	//ask for user input
	fmt.Println("Enter youre first name: ")
	fmt.Scan(&userName)
	fmt.Println("How many tickets do you want?")
	fmt.Scan(&userTickets)

	ticketsLeft = ticketsLeft - int(userTickets)
	bookings = append(bookings, userName)

	fmt.Printf("Thanks for youre input, %v. You have booked %v tickets.\n", userName, userTickets)
	fmt.Printf("There are %v ticktes for the %v left.\n", ticketsLeft, conferenceName)
	fmt.Printf("All the bookings : %v \n", bookings)
	
}
 