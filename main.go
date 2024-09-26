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

	greetUser(conferenceName, conferenceTickets, remainingTicket)

	for {
		//returned values from user input function
		firstName, lastName, email, userTickets := getUserInput()

		//creating a function here for readability
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTicket)

		//improving on the booking ticket logic
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTicket = remainingTicket - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remainingTicketing for %v\n", remainingTicket, conferenceName)
			//Print first names
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTicket == 0 {
				fmt.Printf("Tickets are sold out")
				break
			} else {
				fmt.Printf("%v available\n", remainingTicket)
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name is too short!")

			}

			if !isValidEmail {
				fmt.Println("Your email address requires the @ symbol!")

			}

			if !isValidTicketNumber {
				fmt.Println("Your booking ticket does not meet its requirement!")
			}

		}

	}

}

func greetUser(conferenceName string, conferenceTickets int, remainingTicket uint) {
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTicket, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTicket)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTicket

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}
