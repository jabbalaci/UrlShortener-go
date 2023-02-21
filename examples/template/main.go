package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/jabbalaci/go-urlshortener/lib/jweb"
	"github.com/jabbalaci/go-urlshortener/lib/pronounce"
	"github.com/jabbalaci/go-urlshortener/templates"
)

func main() {
	f, err := os.Create("/tmp/output.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	type Context struct {
		Text  string
		Lines []string
	}

	url := "bit.ly/3oXtVL6"
	id := "3oXtVL6"
	lines := pronounce.Say(id)

	tmpl := template.Must(template.New("html").Parse(templates.ZOOM_HTML))
	context := Context{url, lines}

	var buf bytes.Buffer
	tmpl.Execute(&buf, context)
	result := buf.String()
	f.WriteString(result)
	fmt.Println("# written to", f.Name())
	//
	// open the file in the default browser
	jweb.OpenInBrowser(f.Name())
}
