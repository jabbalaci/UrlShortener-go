package main

import (
	"fmt"
	"strings"
)

func print_help() {
	text := `
urlshortener v{ver}
https://github.com/jabbalaci/UrlShortener-go

An interactive program that shortens a long URL using the Bitly API.
`
	text = strings.TrimSpace(text)
	text = strings.Replace(text, "{ver}", VERSION, 1)

	fmt.Println(text)
}
