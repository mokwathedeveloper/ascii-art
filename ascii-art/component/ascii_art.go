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
