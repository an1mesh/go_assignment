package main

import (
	"fmt"
	"strings"
)

func main() {

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
			break
		}

		fmt.Println("\n")

		// Print placeholder slice
		fmt.Println(placeholder)

		// Print the chances left
		fmt.Printf("chances: %d\n", chances) // render the chances left

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
		fmt.Scanln(&inputStr)
		if len(inputStr) == len(word) && inputStr == word {
			fmt.Println("you won")
			break
		}

		// check for duplicates
		_, duplicate := entries[inputStr]
		if duplicate {
			continue
		}

		// update entries
		entries[inputStr] = true

		isFound := false

		for i, value := range word {
			if inputStr == string(value) {
				placeholder[i] = string(value)
				isFound = true
			}
		}
		if !isFound {
			chances -= 1
		}
	}
}
