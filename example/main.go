package main

import (
	"fmt"
	"github.com/chaseisabelle/icndb"
)

func main() {
	icndb := icndb.New()

	jokes, err := icndb.Jokes("", "")

	if err != nil {
		panic(err)
	}

	fmt.Println("jokes")

	for _, joke := range jokes[:5] {
		fmt.Printf("\t%d: %s\n", joke.Id, joke.Text)
	}

	joke, err := icndb.Joke(jokes[0].Id, "chase", "isabelle")

	if err != nil {
		panic(err)
	}

	fmt.Printf("the jokes on me\n\t%s\n", joke.Text)

	jokes, err = icndb.RandomJokes(5, "our lord and savior", "shrek", make(map[string]bool))

	if err != nil {
		panic(err)
	}

	fmt.Println("random jokes")

	for _, joke := range jokes {
		fmt.Printf("\t%d: %s\n", joke.Id, joke.Text)
	}

	joke, err = icndb.RandomJoke("smash", "mouth")

	if err != nil {
		panic(err)
	}

	fmt.Printf("a random joke\n\t%s\n", joke.Text)

	categories, err := icndb.Categories()

	if err != nil {
		panic(err)
	}

	fmt.Println("categories")

	for _, category := range categories {
		fmt.Printf("\t%s\n", category)
	}
}
