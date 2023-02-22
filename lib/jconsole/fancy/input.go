// Reading from stdin with readline support.
package fancy

import (
	"github.com/chzyer/readline"
	"github.com/jabbalaci/UrlShortener-go/lib/jconsole"
	"github.com/jabbalaci/UrlShortener-go/lib/jos"
)

// Reads a line from stdin using the readline library.
// It also prints a prompt if specified.
// This one doesn't return the error.
func Input(prompt string) string {
	text, err := InputWithError(prompt)
	if err != nil {
		panic(err)
	}
	return text
}

// On Linux, it uses the fancy readline library.
// On Windows I had problems with the readline library in "cmd",
// thus on non-Linux systems we fall back to the normal Input() function.
// I couldn't test it on Mac, so feel free to improve this function.
func InputLinuxOnly(prompt string) string {
	if jos.IsLinux() {
		return Input(prompt)
	} else {
		return jconsole.Input(prompt)
	}
}

// Reads a line from stdin using the readline library.
// It also prints a prompt if specified.
// This one returns the error.
func InputWithError(prompt string) (string, error) {
	line, err := readline.Line(prompt)
	if err != nil {
		return line, err
	}
	// else
	return line, nil
}
