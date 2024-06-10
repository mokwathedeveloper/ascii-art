package main

import (
	"component/asciiart"
	"os"
	"strings"
)

func main() {
	// Determine the correct file based on the arguments
	fileName := asciiart.GetTheCorrectFile(os.Args)

	// Read the contents of the file
	fileContents, err := os.ReadFile(fileName)
	asciiart.HandleError(err)

	// Split the file contents based on the file type
	var lines []string
	if fileName == "thinkertoy.txt" {
		lines = strings.Split(string(fileContents), "\r\n")
	} else {
		lines = strings.Split(string(fileContents), "\n")
	}

	// Create the ASCII graphics map
	asciiMap := asciiart.MakeAsciiGraphicsMap(lines)

	// Output the ASCII graphics for the input string
	asciiart.Output(os.Args[1], lines, asciiMap)
}
