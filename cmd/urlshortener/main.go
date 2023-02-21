package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/jabbalaci/go-urlshortener/lib/bitly"
	"github.com/jabbalaci/go-urlshortener/lib/jconsole"
	"github.com/jabbalaci/go-urlshortener/lib/jweb"
	"github.com/jabbalaci/go-urlshortener/lib/pronounce"
	"github.com/jabbalaci/go-urlshortener/templates"
)

var bold = color.New(color.Bold).PrintlnFunc()
var green = color.New(color.Bold, color.FgGreen).PrintlnFunc()
var red = color.New(color.Bold, color.FgRed).PrintlnFunc()

const OK = "✓"
const NOT_OK = "✗"

func check_api_key() {
	if os.Getenv("BITLY_ACCESS_TOKEN") == "" {
		fmt.Println("Please set the env. variable BITLY_ACCESS_TOKEN")
		os.Exit(1)
	}
}

func copy_to_clipboard(text string) {
	response := jconsole.Input("Copy shortened URL to clipboard [Yn]? ")
	if response == "y" || response == "Y" || response == "" {
		clipboard.WriteAll(text)
		fmt.Print("# copied ")
		green(OK)
	} else {
		fmt.Print("# no ")
		red(NOT_OK)
	}
}

func zoom_into(url string) {
	url = strings.TrimPrefix(url, "https://")
	f, err := os.CreateTemp(os.TempDir(), "urlshortener-*.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	type Context struct {
		Text  string
		Lines []string
	}

	id := strings.Split(url, "/")[1]
	lines := pronounce.Say(id)

	tmpl := template.Must(template.New("html").Parse(templates.ZOOM_HTML))
	context := Context{url, lines}

	var buf bytes.Buffer
	tmpl.Execute(&buf, context)
	result := buf.String()
	f.WriteString(result)
	fmt.Println("# written to", f.Name())
	// open the file in the default browser
	jweb.OpenInBrowser(f.Name())
}

func zoom(url string) {
	response := jconsole.Input("Zoom into this short URL [yN]? ")
	if response == "y" || response == "Y" {
		fmt.Print("# yes ")
		green(OK)
		zoom_into(url)
	} else {
		fmt.Print("# no ")
		red(NOT_OK)
	}
}

func main() {
	check_api_key() // may exit

	original_url := jconsole.Input("Long URL: ")

	short_url, err := bitly.Shorten(original_url)
	if err != nil {
		fmt.Println("Error: the URL couldn't be shortened.")
		os.Exit(1)
	}
	// else
	fmt.Println()
	bold(short_url)

	long_url, err := bitly.Expand(short_url)
	if err != nil {
		fmt.Println("Error: the URL couldn't be expanded.")
		os.Exit(1)
	}
	// else
	long_url = strings.TrimSuffix(long_url, "/")
	fmt.Println()
	fmt.Printf("# expanded from shortened URL: %v", long_url)

	if original_url == long_url {
		green(" (matches)")
	} else {
		red(" (does NOT match)")
		os.Exit(1)
	}

	// if matches
	fmt.Println()
	copy_to_clipboard(short_url)
	zoom(short_url)
}
