package greetings

import (
	"errors"
	"fmt"
	"strings"
)

func GreetWithName(name string) (string, error) {
	displayName := strings.TrimSpace(name)

	if isEmptyString(displayName) {
		return "", errors.New("name argument cannot be left blank")
	}

	message := fmt.Sprintf("Hi %v, welcome to your first Go program!", displayName)

	return message, nil
}

// GreetUser returns a greeting for the named person. If name is missing, uses Guest.
func GreetUser(name string) string {
	displayName := strings.TrimSpace(name)

	if isEmptyString(displayName) {
		displayName = "Guest"
	}

	message := fmt.Sprintf("Hi %v, welcome to your first Go program!", displayName)

	return message
}

func isEmptyString(str string) bool {
	trimmedStr := strings.TrimSpace(str)

	return len(trimmedStr) == 0
}
