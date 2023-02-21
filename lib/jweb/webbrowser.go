package jweb

/*
	based on https://github.com/0x434d53/openinbrowser/blob/master/openinbrowser.go
*/

import (
	"log"
	"os/exec"
	"runtime"
)

// OpenInBrowser opens the specified file in the default browser.
func OpenInBrowser(path string) error {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open", path}
	case "windows":
		args = []string{"cmd", "/c", "start", path}
	default:
		args = []string{"xdg-open", path}
	}
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		log.Printf("openinbrowser: %v\n", err)
		return err
	}
	// else
	return nil
}
