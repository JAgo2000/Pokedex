package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func pepl_loop() {
	scanner := bufio.NewScanner(os.Stdin)
	exit := false
	for {
		fmt.Print(">")
		scanner.Scan()
		text := scanner.Text()
		text_arr := clean_scantext(text)
		if text_arr != nil {
			switch text {
			//case "clear":
			//clear window   fmt.Printf("Echo: %v", text_arr[0])
			case "help":
				fmt.Println("*****************")
				fmt.Println("This is a Repl! ")
				fmt.Println("Valid Commands: ")
				fmt.Println("<help> ")
				fmt.Println("<exit> ")
				fmt.Println("*****************")
			case "exit":
				fmt.Println("exit now")
				exit = true
			default:
				fmt.Printf("<%v> is a unknown comand, type <help> for instructions", text_arr[0])
				fmt.Println()
			}
			if exit {
				break
			}
		}
	}
}

func clean_scantext(text string) []string {
	if len(text) == 0 {
		return nil
	}
	text_lower := strings.ToLower(text)
	arr_str := strings.Fields(text_lower)
	return arr_str
}
