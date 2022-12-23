package main

import (
	"fmt"
	"os"
)

func ErrorHandled(err error) {
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
}
