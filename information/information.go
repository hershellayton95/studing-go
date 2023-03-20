package information

import "fmt"

func SetGetInformation() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("What's your first name?")
	fmt.Scan(&firstName)
	fmt.Println("What's your last name?")
	fmt.Scan(&lastName)
	fmt.Println("What's your email?")
	fmt.Scan(&email)
	fmt.Println("How many tickets would you order?")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
