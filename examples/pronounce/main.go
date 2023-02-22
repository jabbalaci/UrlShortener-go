package main

import (
	"fmt"

	"github.com/jabbalaci/UrlShortener-go/lib/pronounce"
)

func main() {
	id := "3oXtVL6"

	fmt.Println(id)
	fmt.Println()

	say := pronounce.Say(id)
	for _, word := range say {
		fmt.Println(word)
	}
}
