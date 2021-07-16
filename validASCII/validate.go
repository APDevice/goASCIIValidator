package validASCII

type extended bool

const (
	Extended      extended = true
	stdASCII      rune     = 127
	extendedASCII rune     = 255
)

var asciiSet rune

// Validate takes a string and returns false if non-ASCII characters are found, otherwise returns true
func Validate(str string, ext ...extended) bool {
	if ext[0] {
		asciiSet = extendedASCII
	} else {
		asciiSet = stdASCII
	}

	for _, char := range str {
		if char > asciiSet {
			return false
		}
	}

	return true
}

// CountNonASCII takes a string and counts the number of non-ASCII characters in it
func CountNonASCII(str string, ext ...extended) int {
	if ext[0] {
		asciiSet = extendedASCII
	} else {
		asciiSet = stdASCII
	}

	var cnt int
	for _, char := range str {
		if char > asciiSet {
			cnt++
		}
	}

	return cnt
}

// CountNonASCII takes a string and counts the number of non-ASCII characters in it
func CountASCII(str string, ext ...extended) int {
	if ext[0] {
		asciiSet = extendedASCII
	} else {
		asciiSet = stdASCII
	}

	var cnt int
	for _, char := range str {
		if char <= asciiSet {
			cnt++
		}
	}

	return cnt
}
