package utils

import (
	"math/rand/v2"
	"strings"
)

func GeneratePassword(size int) string {
	var generatedPassword strings.Builder
	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharacters := "!@#$%^&*()_+-=[]{}|;:'\",.<>?/`~"
	pinCode := "1234567890"

	defaultCriteria := alphabetLower + alphabetUpper + specialCharacters + pinCode

	for i := 0; i < size; i++ {
		randomNumber := rand.IntN(len(defaultCriteria))
		generatedCharacter := string(defaultCriteria[randomNumber])
		generatedPassword.WriteString(generatedCharacter)
	}

	return generatedPassword.String()
}
