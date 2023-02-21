package jweb

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// based on https://terminalroot.com/how-to-open-url-in-default-browser-in-go-python-ruby-and-rust/
func OpenInBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Printf("OpenInBrowser: %v\n", err)
		return err
	}
	// else
	return nil
}
