package main

import (
	"fmt"
	"os"

	"github.com/tetzng/typolice"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: typolice [TEXT]")
		os.Exit(1)
	}

	text := os.Args[1]

	results := typolice.Run(text)

	for _, result := range results {
		fmt.Printf("Typo: %s, Preferred: %s\n", result.Typo, result.Preferred)
	}

	if len(results) == 0 {
		fmt.Println("No typos found.")
	} else {
		fmt.Println("Total typos found: ", len(results))
	}
}
