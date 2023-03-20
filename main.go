package main

import (
	"GoTest/information"
	"fmt"
	"sync"
	"time"
)

var conferanceName string = "Go Conferance"

const conferanceTicket int = 50

var remainsTickets uint = 50

var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetings(conferanceName, conferanceTicket, remainsTickets)

	// for { // senza niente è un look infinito
	firstName, lastName, email, userTickets := information.SetGetInformation()

	isValidName, isValidEmail, isValidTicket := authentication(firstName, lastName, email, userTickets, remainsTickets)

	if isValidName && isValidEmail && isValidTicket {
		remainsTickets = remainsTickets - userTickets

		var userData = UserData{
			firstName:   firstName,
			lastName:    lastName,
			email:       email,
			userTickets: userTickets,
		}

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		bookings = append(bookings, userData)
		fmt.Printf("List of booking %v\n", bookings)

		fmt.Printf("Thank you %v %v to have orderd %v tickets. You'll receve a confirm email on %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remains for %v\n", remainsTickets, conferanceName)

		firstNames := getFirstName(bookings)
		fmt.Printf("all first name %v\n", firstNames)

		if remainsTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
		}
	} else {
		if !isValidName {
			fmt.Println("The name is too short")
		}
		if !isValidEmail {
			fmt.Println("The email is to short or it doesn't contain @")
		}
		if !isValidTicket {
			fmt.Println("Number of tickets are invalid")
		}
	}
	wg.Wait()
}

// func main() {

// 	greetings(conferanceName, conferanceTicket, remainsTickets)

// 	// for { // senza niente è un look infinito
// 	for remainsTickets > 0 && len(bookings) < 50 {
// 		firstName, lastName, email, userTickets := information.SetGetInformation()

// 		isValidName, isValidEmail, isValidTicket := authentication(firstName, lastName, email, userTickets, remainsTickets)

// 		if isValidName && isValidEmail && isValidTicket {
// 			remainsTickets = remainsTickets - userTickets

// 			// var userData = make(map[string]string)
// 			// userData["firstName"] = firstName
// 			// userData["lastName"] = lastName
// 			// userData["email"] = email
// 			// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

// 			// userData["firstName"] = firstName
// 			// userData["lastName"] = lastName
// 			// userData["email"] = email
// 			// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

// 			// bookings = append(bookings, firstName+" "+lastName)
// 			var userData = UserData{
// 				firstName:   firstName,
// 				lastName:    lastName,
// 				email:       email,
// 				userTickets: userTickets,
// 			}

// 			go sendTicket(userTickets, firstName, lastName, email)
// 			bookings = append(bookings, userData)
// 			fmt.Printf("List of booking %v\n", bookings)

// 			fmt.Printf("Thank you %v %v to have orderd %v tickets. You'll receve a confirm email on %v\n", firstName, lastName, userTickets, email)
// 			fmt.Printf("%v tickets remains for %v\n", remainsTickets, conferanceName)

// 			// fmt.Printf("Array of all user who's booked %v\n", booking)
// 			// fmt.Printf("Type of array %T\n", booking)
// 			// fmt.Printf("len of array %v\n", len(booking))

// 			// fmt.Printf("Slice of all user who's booked %v\n", booking)
// 			// fmt.Printf("Type of slice %T\n", booking)
// 			// fmt.Printf("len of slice %v\n", len(booking))

// 			firstNames := getFirstName(bookings)
// 			fmt.Printf("all first name %v\n", firstNames)

// 			if remainsTickets == 0 {
// 				fmt.Println("Our conference is booked out. Come back next year.")
// 				break
// 			}
// 		} else {
// 			if !isValidName {
// 				fmt.Println("The name is too short")
// 			}
// 			if !isValidEmail {
// 				fmt.Println("The email is to short or it doesn't contain @")
// 			}
// 			if !isValidTicket {
// 				fmt.Println("Number of tickets are invalid")
// 			}
// 		}
// 	}
// }

func greetings(conferanceName string, conferanceTicket int, remainsTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferanceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferanceTicket, remainsTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

// func getFirstName(bookings []map[string]string) []string {
// 	firstNames := []string{}
// 	for _, booking := range bookings {
// 		firstNames = append(firstNames, booking["firstName"])
// 	}
// 	return firstNames
// }

func getFirstName(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	fmt.Println("##################")
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("##################")
	wg.Done()
}
