package component

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// ASCII constants
const (
	firstAscii = 32
	lastAscii  = 127
)

// MakeAsciiGraphicsMap maps runes to their corresponding starting points in the ASCII graphic.
func MakeAsciiGraphicsMap(fileContents []string) map[rune]int {
	asciiMap := map[rune]int{}
	startAscii := 32
	for i, content := range fileContents {
		if len(content) == 0 || len(content) == 1 {
			asciiMap[rune(startAscii)] = i + 1
			startAscii++
		}
	}
	return asciiMap
}

// GetTheCorrectFile determines which ASCII graphic file to use based on the flag passed.
func GetTheCorrectFile(args []string) string {
	if len(args) == 1 {
		fmt.Println("Not enough arguments. To run the program type go run . <text> [-flag] where flag can be -s or -t")
		os.Exit(0)
	}

	if len(args) > 3 {
		fmt.Println("Too many arguments. To run the program type go run . <text> [-flag] where flag can be -s or -t")
		os.Exit(0)
	}

	if len(args) == 3 {
		switch args[2] {
		case "-t", "-thinkertoy", "thinkertoy":
			return "thinkertoy.txt"
		case "-s", "-shadow", "shadow":
			return "shadow.txt"
		case "-standard", "standard":
			return "standard.txt"
		default:
			fmt.Println("Invalid flag. To run the program type go run . <text> [-flag] where flag can be -s or -t")
			os.Exit(0)
		}
	}
	return "standard.txt"
}

// Output prints the ASCII graphics for a given string, using the file contents and ASCII map provided.
func Output(str string, fileContents []string, asciiMap map[rune]int) {
	if len(str) == 0 {
		return
	}

	// Replace escape sequences with their actual meanings
	str = strings.ReplaceAll(str, "\\n", "\n")
	str = strings.ReplaceAll(str, "\\t", "\t")
	str = strings.ReplaceAll(str, "\\r", "\r")

	// Split the input into lines
	lines := strings.Split(str, "\n")

	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range line {
				if char < rune(firstAscii) || char > rune(lastAscii) {
					if unicode.IsSpace(char) {
						// Handle spaces and other white spaces correctly
						fmt.Print(strings.Repeat(" ", 8))
					} else {
						fmt.Println("Unprintable ASCII character(s) in input. To run the program type go run . <text(printable ascii)> [-flag] where flag can be -s or -t")
						os.Exit(0)
					}
				} else {
					if ascii, ok := asciiMap[char]; ok {
						fmt.Print(fileContents[ascii+i])
					} else {
						fmt.Print(strings.Repeat(" ", 8)) // Handle characters not in the map with spaces
					}
				}
			}
			fmt.Println()
		}
	}
}

// HandleError outputs an error message and exits the program.
func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
