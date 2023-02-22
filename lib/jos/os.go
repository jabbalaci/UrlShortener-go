package jos

import (
	"runtime"
)

// Returns the name of the operating system.
func GetOperatingSystem() string {
	// To see all possible combinations of GOOS and GOARCH, run the command
	// $ go tool dist list
	return runtime.GOOS
}

// Are we on Linux?
func IsLinux() bool {
	return GetOperatingSystem() == "linux"
}

// Are we on Windows?
func IsWindows() bool {
	return GetOperatingSystem() == "windows"
}

// Are we on Mac?
func IsMac() bool {
	return GetOperatingSystem() == "darwin"
}

// Number of logical CPUs usable by the current process.
func NumCPU() int {
	return runtime.NumCPU()
}
