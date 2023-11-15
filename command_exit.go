package main

import (
	"fmt"
	"os"
)

func callbackexit() error {
	fmt.Println("Exiting Now")
	os.Exit(0)
	return nil
}
