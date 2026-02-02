package main

import (
	// "booking_app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Golang"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}
var mutex = sync.Mutex{}

func main() {

	greetUsers()

	firstName, lastName, emailAddress, userTickets := getUserInputs()
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInputs(firstName, lastName, emailAddress, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTickets(userTickets, firstName, lastName, emailAddress)

		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, emailAddress)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v \n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("firstname or lastname entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address doesn't contain @ symbol")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to the %v conference\n", conferenceName)
	fmt.Printf("We have a total of %v ickets and %v are still available\n", conferenceTickets, remainingTickets)
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&emailAddress)
	fmt.Println("Enter how many tickets you want buy: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailAddress, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, emailAddress string) {
	// mutex.Lock()
	// defer mutex.Unlock()

	remainingTickets = remainingTickets - userTickets
	var userData = userData{
		firstName:   firstName,
		lastName:    lastName,
		email:       emailAddress,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List if bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation booking at %v \n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for: %v \n", remainingTickets, conferenceName)

}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTickets(userTickets uint, firstName string, lastName string, emailAddress string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket: \n %v to email address %v\n", ticket, emailAddress)
	fmt.Println("####################")
	wg.Done()
}
