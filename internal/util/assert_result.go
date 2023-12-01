package util

import "fmt"

func AssertResult(name string, result string, expected string) {
	if result != expected {
		fmt.Printf("[%-20s] Expected %s but got %s instead.\n", name, expected, result)
	} else {
		fmt.Printf("[%-20s] Result: %s\n", name, result)
	}
}
