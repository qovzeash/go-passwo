package utils

import (
	"math/rand/v2"
	"strings"
)

func GeneratePassword(size int, criteria string) string {
	var generatedPassword strings.Builder
	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharacters := "!@#$%^&*()_+-=[]{}|;:'\",.<>?/`~"
	pinCode := "1234567890"

	switch criteria {
	case "defaultCriteria":
		criteria = alphabetLower + alphabetUpper + specialCharacters + pinCode
	case "noSpecial":
		criteria = alphabetLower + alphabetUpper + pinCode
	case "pinCode":
		criteria = pinCode
	}

	for i := 0; i < size; i++ {
		randomNumber := rand.IntN(len(criteria))
		generatedCharacter := string(criteria[randomNumber])
		generatedPassword.WriteString(generatedCharacter)
	}

	return generatedPassword.String()
}
