package main

import (
	"bufio"
	"fmt"
	"os"
)

func pepl_loop() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >")
		scanner.Scan()
		text := scanner.Text()
		fmt.Printf("Echo: %v", text)
		fmt.Println()
	}
}
