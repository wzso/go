package main

import (
	"fmt"
	"runtime"
	"time"
)

func osInfo() {
	s := "Go runs on "
	switch os := runtime.GOOS; os {
	case "darwin":
		s += "macOS."
	case "Linux":
		s += "Linux"
	default:
		s += os
	}

	fmt.Println(s)
}



func weekday() {
	today := time.Now().Weekday()
	fmt.Println(today)
}

func main() {
	osInfo()
	weekday()
}
