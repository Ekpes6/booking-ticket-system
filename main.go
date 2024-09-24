package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTicket uint = 50
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTicket, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTicket)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your First Name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last Name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email Name: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTicket {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTicket, userTickets)
			break
		}

		remainingTicket = remainingTicket - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remainingTicketing for %v\n", remainingTicket, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		if remainingTicket == 0 {
			fmt.Printf("Tickets are sold out")
			break
		} else {
			fmt.Printf("%v available\n", remainingTicket)
		}
	}

}
