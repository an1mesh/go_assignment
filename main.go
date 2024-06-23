package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("go basics")

	word := "panda"

	// Look for entries made by the user.(Input is taken and type is inferred using :=)
	entries := map[string]bool{}

	// List of "_" (blanks) corrosponding to word length.([]string{} is used to declare empty slice)
	placeholder := []string{}

	// No. of chances a user gets
	chances := 5

	// Loop to add guessed word in placeholder slice.
	for i := 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")
	}

	for {
		// If user enters a wrong letter or the wrong word, they lose a chance.
		userInput := strings.Join(placeholder, "")
		if chances == 0 && userInput != word {
			fmt.Println("game over, try again")
			break
		}
		// User won
		if userInput == word {
			fmt.Println("you win")
		}

		fmt.Println("\n")

		// Print placeholder slice
		fmt.Println(placeholder)

		keys := []string{}

		// If we are not using a variable agian, declare with _
		for k, _ := range entries {

			// appends key of each entry of map entries
			keys = append(keys, k)
		}
		fmt.Println(keys)
		fmt.Printf("Guess a letter or the word: ")

		// Take input
		inputStr := ""
		// address of inputStr is passed so there is no copy of inputStr is made and its passed by reference
		fmt.Scanln(&inputStr)

		// Compare and update entries, placeholder and chances.
		entries[inputStr] = true
	}
}
