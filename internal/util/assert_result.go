package util

import (
	"fmt"
	"os"
)

func AssertResult[R comparable](name string, result R, expected R) {
	if result != expected {
		fmt.Printf("[%-15s] Expected %+v but got %+v instead.\n", name, expected, result)
		fmt.Println("Test failed! Exiting...")
		os.Exit(1)
	} else {
		fmt.Printf("[%-15s] Result: %+v\n", name, result)
	}
}
