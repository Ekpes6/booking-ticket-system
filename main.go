package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTicket uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUser()

	for {
		//returned values from user input function
		firstName, lastName, email, userTickets := getUserInput()

		//creating a function here for readability
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTicket)

		//improving on the booking ticket logic
		if isValidName && isValidEmail && isValidTicketNumber {
			//booking ticket function
			bookTicket(userTickets, firstName, lastName, email)

			//Print first names
			firstNames := getFirstNames()
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

// greet or welcome message
func greetUser() {
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTicket, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTicket)
	fmt.Println("Get your tickets here to attend")
	//fmt.Println(helper.MyScope)

}

// get all first names and add to a slice object
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

// get user input
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTickets

	//create a user data map
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remainingTicketing for %v\n", remainingTicket, conferenceName)
	fmt.Println(bookings)

}
