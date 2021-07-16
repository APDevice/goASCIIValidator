package ASCIIValidator

var (
	asciiSet int
)

func init() {
	asciiSet = 127
}

// Extended expands valid characters to the 8-bit extended ASCII set if true. Resets to standard 7-bit ASCII set if false.
func Extended(isExtended bool) {
	if isExtended {
		asciiSet = 255
	} else {
		asciiSet = 127
	}
}

// Validate takes a string and returns FALSE if non-ASCII characters are found, otherwise returns true
func Validate(str string) bool {
	for _, char := range str {
		if char > asciiSet {
			return false
		}
	}

	return true
}
