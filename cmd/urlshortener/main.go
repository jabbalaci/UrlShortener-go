package main

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/jabbalaci/UrlShortener-go/lib/bitly"
	"github.com/jabbalaci/UrlShortener-go/lib/jconsole/fancy"
	"github.com/jabbalaci/UrlShortener-go/lib/jweb"
	"github.com/jabbalaci/UrlShortener-go/lib/pronounce"
	"github.com/jabbalaci/UrlShortener-go/templates"
	qrcode "github.com/skip2/go-qrcode"
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
	response := fancy.InputLinuxOnly("Copy shortened URL to clipboard [Yn]? ")
	if response == "y" || response == "Y" || response == "" {
		clipboard.WriteAll(text)
		fmt.Print("# copied ")
		green(OK)
	} else {
		fmt.Print("# no ")
		red(NOT_OK)
	}
}

func createQrCode(basename, url string) (string, error) {
	pngFile := strings.TrimSuffix(basename, ".html") + ".png"
	err := qrcode.WriteFile(url, qrcode.Medium, 512, pngFile) // 512x512 pixels
	if err != nil {
		log.Printf("createQrCode: %v\n", err)
		return "", err
	}
	return pngFile, nil
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
		Png   string
	}

	id := strings.Split(url, "/")[1]
	lines := pronounce.Say(id)
	image_name, _ := createQrCode(f.Name(), url)

	tmpl := template.Must(template.New("html").Parse(templates.ZOOM_HTML))
	context := Context{url, lines, filepath.Base(image_name)}

	var buf bytes.Buffer
	tmpl.Execute(&buf, context)
	result := buf.String()
	f.WriteString(result)
	fmt.Println("# written to", f.Name())
	// open the file in the default browser
	jweb.OpenInBrowser(f.Name())
}

func zoom(url string) {
	response := fancy.InputLinuxOnly("Zoom into this short URL [yN]? ")
	if response == "y" || response == "Y" {
		fmt.Print("# yes ")
		green(OK)
		zoom_into(url)
	} else {
		fmt.Print("# no ")
		red(NOT_OK)
	}
}

func match(original_url, expanded_url string) bool {
	if original_url == expanded_url {
		return true
	}
	// else
	decoded_url, err := url.QueryUnescape(original_url)
	if err != nil {
		return false
	}
	decoded_url = strings.ReplaceAll(decoded_url, " ", "+")
	return decoded_url == expanded_url
}

func main() {
	arg := ""
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	if arg == "-h" || arg == "--help" {
		print_help()
		os.Exit(0)
	}

	check_api_key() // may exit

	original_url := fancy.InputLinuxOnly("Long URL: ")

	short_url, err := bitly.Shorten(original_url)
	if err != nil {
		fmt.Println("Error: the URL couldn't be shortened.")
		os.Exit(1)
	}
	// else
	fmt.Println()
	bold(short_url)

	expanded_url, err := bitly.Expand(short_url)
	if err != nil {
		fmt.Println("Error: the URL couldn't be expanded.")
		os.Exit(1)
	}
	// else
	expanded_url = strings.TrimSuffix(expanded_url, "/")
	fmt.Println()
	fmt.Printf("# expanded from shortened URL: %v", expanded_url)

	if match(original_url, expanded_url) {
		green(" (matches)")
	} else {
		red(" (does NOT match)")
		os.Exit(1)
	}

	// if they match
	fmt.Println()
	copy_to_clipboard(short_url)
	zoom(short_url)
}
