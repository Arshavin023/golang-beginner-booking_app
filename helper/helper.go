package helper

import (
	"strings"
	"sync"
	"unicode"
)

var mutex = sync.Mutex{}

func ValidateUserInputs(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 && isAlphabetic(firstName) && isAlphabetic(lastName)
	isValidEmail := strings.Contains(email, "@")

	mutex.Lock()
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	mutex.Unlock()

	return isValidName, isValidEmail, isValidTicketNumber
}

// isAlphabetic returns true only if every rune in s is a letter.
// Rejects names like "12", "43", or "Uche1".
func isAlphabetic(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
