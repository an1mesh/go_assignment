package main

import (
	"fmt"
)

// struct of Hangman
type Hangman struct {
	Name        string
	Entries     map[string]bool
	Placeholder []string
	Chances     int
}

// struct of Game
type Game struct {
	Name string
	Hangman
}

// function
func play(h Hangman, g *Game) (won bool, err error) {
	h.Name = "Hangman 2.0"
	h.Placeholder = append(h.Placeholder, "c")
	h.Entries["lion"] = true

	g.Name = "Animesh Sinha"
	g.Placeholder = append(g.Placeholder, "d")
	return
}

func main() {
	fmt.Println("Structs and Functions")
	h := Hangman{}
	h.Entries = make(map[string]bool)
	h.Placeholder = []string{}

	// Object of Hangman struct
	h1 := Hangman{
		Entries:     make(map[string]bool),
		Placeholder: []string{"a", "b"},
		Chances:     8,
		Name:        "Animesh",
	}
	h1.Entries["deer"] = true

	// Object of Game struct
	g1 := Game{
		Name: "Animesh",
		Hangman: Hangman{
			Placeholder: []string{"a", "b"},
			Entries:     make(map[string]bool),
			Chances:     8,
		},
	}

	won, err := play(h1, &g1)

	if err != nil {
		panic(err)
	}

	fmt.Println(won)

	fmt.Println(g1, h1)
}
