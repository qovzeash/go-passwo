package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	passwordSize := 21
	var generatedPassword strings.Builder

	for i := 0; i < passwordSize; i++ {
		number := rand.IntN(len(alphabet))
		generatedLetter := string(alphabet[(number)])
		generatedPassword.WriteString(generatedLetter)
	}

	fmt.Println(generatedPassword.String())
}
