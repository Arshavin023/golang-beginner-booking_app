package main

import (
	"strings"
)

func ValidateUserInputs(firstName string, lastName string, emailAddress string,
	userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(emailAddress, "@")
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
