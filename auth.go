package main

import "strings"

func authentication(firstName string, lastName string, email string, userTickets uint, remainsTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainsTickets
	return isValidName, isValidEmail, isValidTicket
}
