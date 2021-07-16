package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/APDevice/goASCIIValidator/validASCII"
)

/* checkStrings runs through all inputed strings, checks their validity, and outputs result

Parameters:
	- array of strings to parse
	- charRange representing the ASCII standard to check against

Output:
	- one line for each string, in the format "String %: True/False"
*/
func checkStrings(strs []string, charRange validASCII.CharRange) {
	var isValid bool
	for i, str := range strs {
		isValid = validASCII.Validate(str, charRange)
		fmt.Printf("String %d: %t\n", i+1, isValid)
	}
}

/* checkStringsConcise runs through all inputed strings, checks their validity, and outputs result in a concise stream

Parameters:
	- array of strings to parse
	- charRange representing the ASCII standard to check against

Output:
	- single line of characters, with 1 representing a valid string, 0 representing an invalid one
*/
func checkStringsConcise(strs []string, charRange validASCII.CharRange) {
	var isValid bool
	for _, str := range strs {
		isValid = validASCII.Validate(str, charRange)
		if isValid {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
	fmt.Print("\n")
}
func main() {
	checkExtended := flag.Bool("e", false, "check string against the extended ASCII range")
	printConcise := flag.Bool("c", false, "generate a concise ouput for each string of 1 for valid, 0 for invalid")
	flag.Parse()
	input := flag.Args()

	// check for insufficient arguments
	if len(input) == 0 {
		fmt.Fprintf(os.Stderr, "Insufficient arguments: minimum one string needed to parse")
		os.Exit(1)
	}

	if *printConcise {
		if *checkExtended {
			checkStringsConcise(input, validASCII.Extended)
		} else {
			checkStringsConcise(input, validASCII.Standard)
		}
	} else {
		if *checkExtended {
			checkStrings(input, validASCII.Extended)
		} else {
			checkStrings(input, validASCII.Standard)
		}
	}
}
