package main

import (
	"os"
	"strings"

	"ascii-art/component"
)

func main() {
	// Determine the correct file based on the arguments
	fileName := component.GetTheCorrectFile(os.Args)

	// Validate the file checksum
	err := ValidateFileChecksum(fileName)
	component.HandleError(err)

	// Read the contents of the file
	fileContents, err := os.ReadFile(fileName)
	component.HandleError(err)

	// Split the file contents based on the file type
	var lines []string
	if fileName == "thinkertoy.txt" {
		lines = strings.Split(string(fileContents), "\r\n")
	} else {
		lines = strings.Split(string(fileContents), "\n")
	}

	// Create the ASCII graphics map
	asciiMap := component.MakeAsciiGraphicsMap(lines)

	// Output the ASCII graphics for the input string
	component.Output(os.Args[1], lines, asciiMap)
}
