package component

// ASCII constants
const (
	firstAscii = 32
	lastAscii  = 127
)
// MakeAsciiGraphicsMap maps runes to their corresponding starting points in the ASCII graphic.
func MakeAsciiGraphicsMap(fileContents []string) map[rune]int {
	asciiMap := map[rune]int{}
	for i, content := range fileContents {
		if len(content) == 0 || len(content) == 1 {
			asciiMap[rune(firstAscii+i)] = i + 1
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
		case "-s", "-shadow":
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
