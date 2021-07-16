/* validASCII includes a number of tools for comparing imput to the ASCII standard.
This library was made for the purpose of checking whether a URL or email address has been
spoofed using a common trick of replacing english characters with similar-looking unicode ones */
package validASCII

type charRange rune

const (
	// includes all English characters
	Standard charRange = 127
	// Includes accented characters used in many European languages (also known as latin-1 suplimemt of the unicode standard)
	Extended charRange = 255
)

/* Validate checks whether all characters in a string are valid ASCII characters

Parameters:
	- String to be validated
	- charRange representing the ASCII standard to check against (validASCII.Standard for 7-bit, validASCII.Extended for 8-bit)
Output:
	- boolean value (true if all characters are ASCII, else false)
*/
func Validate(str string, asciiSet charRange) bool {
	for _, char := range str {
		if char > rune(asciiSet) {
			return false
		}
	}

	return true
}

/* CountNonASCII counts the number of non-ASCII characters in a string

Parameters:
	- String to be scanned
	- charRange representing the ASCII standard to check against (validASCII.Standard for 7-bit, validASCII.Extended for 8-bit)
Output:
	- int representing the total number of non-ASCII characters
*/
func CountNonASCII(str string, asciiSet charRange) int {
	var cnt int
	for _, char := range str {
		if char > rune(asciiSet) {
			cnt++
		}
	}

	return cnt
}

/* CountASCII counts the total number of valid ASCII characters in a string

Parameters:
	- String to be scanned
	- charRange representing the ASCII standard to check against (validASCII.Standard for 7-bit, validASCII.Extended for 8-bit)
Output:
	- int representing the total number of ASCII characters
*/
func CountASCII(str string, asciiSet charRange) int {

	var cnt int
	for _, char := range str {
		if char <= rune(asciiSet) {
			cnt++
		}
	}

	return cnt
}
