/*
	Project: validASCII
	Author: Dylan Luttrell
	License: GNU GPLv3
	Description: validASCII includes a number of tools for comparing imput to the ASCII standard.
	This library was made for the purpose of checking whether a URL or email address has been
	spoofed using a common trick of replacing latin characters with similar-looking unicode ones
*/
package validASCII

import "strings"

type CharRange rune

const (
	// includes all basic latin characters and other symbols commonly used in English
	Standard CharRange = 127
	// Includes accented characters used in many European languages (aka latin-1 supplement of the unicode standard)
	Extended CharRange = 255
)

/* Validate checks whether all characters in a string are valid ASCII characters

Parameters:
	- String to be validated
	- charRange representing the ASCII standard to check against (validASCII.Standard for 7-bit, validASCII.Extended for 8-bit)
Output:
	- boolean value (true if all characters are ASCII, else false)
*/
func Validate(str string, asciiSet CharRange) bool {
	for _, char := range str {
		if char > rune(asciiSet) {
			return false
		}
	}

	return true
}

/* Mark checks whether all characters in a string are valid ASCII characters

Parameters:
	- String to be validated
	- charRange representing the ASCII standard to check against (validASCII.Standard for 7-bit, validASCII.Extended for 8-bit)
Output:
	- boolean value (true if all characters are ASCII, else false)
*/
func Mark(str string, asciiSet CharRange) string {
	var sb strings.Builder
	sb.Grow(len(str))
	markChar := '�'

	for _, char := range str {
		if char <= rune(asciiSet) {
			sb.WriteRune(char)
		} else {
			sb.WriteRune(markChar)
		}
	}

	return sb.String()
}

/* CountNonASCII counts the number of non-ASCII characters in a string

Parameters:
	- String to be scanned
	- charRange representing the ASCII standard to check against (validASCII.Standard for 7-bit, validASCII.Extended for 8-bit)
Output:
	- int representing the total number of non-ASCII characters
*/
func CountNonASCII(str string, asciiSet CharRange) int {
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
func CountASCII(str string, asciiSet CharRange) int {

	var cnt int
	for _, char := range str {
		if char <= rune(asciiSet) {
			cnt++
		}
	}

	return cnt
}
