package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	fonts := ""

	const usage = "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard"
	if len(args) < 1 || len(args) > 3 {
		fmt.Println(usage)
		return
	}

	// chose by default font when not chosen 
	if len(args) == 2 {
		fonts = "fonts/standard.txt"
	} else {
		switch args[2] {

		case "standard", "standard.txt":
			fonts = "fonts/standard.txt"

		case "shadow", "shadow.txt":
			fonts = "fonts/shadow.txt"

		case "thinkertoy", "thinkertoy.txt":
			fonts = "fonts/thinkertoy.txt"
		}
	}

	// check the new line and if there is nothing at the beginning of the input
	if args[1] == "\\n" {
		fmt.Printf("\n")
	} else if args[1] != "" {
		fontMap(args[1], fonts)
	}
}

func fontMap(input, font string) {
	// read the font and map it to ascii
	readFile, err := os.ReadFile(font)
	if err != nil {
		fmt.Printf("Could not read the content in the file due to %v", err)
	}
	slice := strings.Split(string(readFile), "\n")

	ascii := make(map[int][]string)
	i := 31

	for _, ch := range slice {
		if ch == "" {
			i++
		} else {
			ascii[i] = append(ascii[i], ch)
		}
	}

	// handle user input and map it to ascii
	lines := strings.Split(input, "\\n")

	for _, line := range lines {
		characters := []rune(line)

		if line != "" {
			for i := 0; i < 8; i++ {

				output := ""
				for _, ch := range characters {
					if ch >= 32 && ch <= 126 {
						output = output + ascii[int(ch)][i]
					} else {
						fmt.Print("Input includes non ascii character(s)\nPlease use ascii character(s)\n")
						os.Exit(0)
					}
				}
				fmt.Print(output + "\n")

			}
		} else {
			fmt.Printf("\n")
		}

	}
}
