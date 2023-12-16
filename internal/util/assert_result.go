package util

import (
	"fmt"
	"os"
)

func AssertResult(name string, result string, expected string) {
	if result != expected {
		fmt.Printf("[%-20s] Expected %s but got %s instead.\n", name, expected, result)
		fmt.Println("Test failed! Exiting...")
		os.Exit(1)
	} else {
		fmt.Printf("[%-20s] Result: %s\n", name, result)
	}
}
