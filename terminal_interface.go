/*
	Project: asciiValidator
	Author: Dylan Luttrell
	License: GNU GPLv3
	Description: Provides an interface for the validASCII library.
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"./validASCII"
)

// checkErr exits with an error if error occurs
func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

/* _checkStringExtended checks string and returns a string representing the validity of it

Parameters:
	- str: string to parse
	- charRange: validASCII.CharRange representing the ASCII standard to check against

Return:
	- "valid" if code is pure ASCII, else count of invalid characters
*/
func _checkStringExtended(str string, charRange validASCII.CharRange) string {
	if validASCII.Validate(str, charRange) {
		return "valid\n"
	}
	return fmt.Sprintf("%d invalid characters\n", validASCII.CountNonASCII(str, charRange))
}

/* _checkStringConcise checks string and returns a byte representing the validity of it

Parameters:
	- array of strings to parse
	- charRange representing the ASCII standard to check against

Return:
	- single byte, with 1 representing a valid string, 0 representing an invalid one
*/
func _checkStringsConcise(str string, charRange validASCII.CharRange) byte {
	if validASCII.Validate(str, charRange) {
		return '1'
	}
	return '0'
}

/* checkStringsFromStdIN: runs through all inputed strings, checks their validity, and outputs result

Parameters:
	- strs: string array of strings to parse
	- charRange: validASCII.CharRange representing the ASCII standard to check against
	- consise: a bool that concise (true) and extended (false) output
	- fileOut: string - filename for output file. If left blank, will output to STDOUT.

Output (file/STDOUT):
	- results of string check
*/
func checkStringsFromStdIN(strs []string, charRange validASCII.CharRange, concise bool) string {
	var (
		result strings.Builder
	)
	if concise {
		for _, str := range strs {
			result.WriteByte(_checkStringsConcise(str, charRange))
		}
		result.WriteByte('\n')
	} else {
		for _, str := range strs {
			result.WriteString(_checkStringExtended(str, charRange))
		}
	}

	return result.String()
}

/* checkFile: runs through all characters in text file, checking their validity, and outputs result

Parameters:
	- fileIn: string - filename of input file
	- charRange: validASCII.CharRange representing the ASCII standard to check against
	- consise: a bool that concise (true) and extended (false) output
	- fileOut: string - filename for output file. If left blank, will output to STDOUT.

Output (file/STDOUT):
	- results of string check
*/
func checkFile(fileIn string, charRange validASCII.CharRange, concise bool) string {
	var (
		data   []byte
		err    error
		result string
	)
	data, err = os.ReadFile(fileIn)
	checkErr(err)

	if concise {
		result = string(_checkStringsConcise(string(data), charRange)) + "\n"
	} else {
		result = _checkStringExtended(string(data), charRange)
	}

	return result
}

/* outputResults prints result string to file

Parameters:
	- fileOut(string): name of file to output to
	- result(string): string to be outputed

*/
func outputResults(fileOut string, result string) {
	var (
		err  error
		fout *os.File
	)

	if len(fileOut) > 0 {
		fout, err = os.Create(fileOut)
		checkErr(err)
	}

	_, err = fout.WriteString(result)
	checkErr(err)

	fout.Close()
}

// Main function
func main() {
	var (
		checkExtendedFlag = flag.Bool("e", false, "check string against the extended ASCII range")
		printConciseFlag  = flag.Bool("c", false, "generate a concise ouput for each string of 1 for valid, 0 for invalid")
		fileIn            = flag.String("r", "", "read from file")
		fileOut           = flag.String("w", "", "write results to file")
		result            string // store results of validator
	)
	flag.Parse()

	// if no strings of file given, exit with error
	if len(flag.Args()) == 0 && len(*fileIn) == 0 {
		fmt.Fprintf(os.Stderr, "Insufficient arguments: minimum one string needed to parse\n")
		os.Exit(1)
	}

	// set ascii range, defaulting to 128 unless extended flag is given
	var asciiSet validASCII.CharRange
	if *checkExtendedFlag {
		asciiSet = validASCII.Extended
	} else {
		asciiSet = validASCII.Standard
	}

	// check file if given, else check STDIN
	if len(*fileIn) > 0 {
		result = checkFile(*fileIn, asciiSet, *printConciseFlag)
	} else {
		stdin := flag.Args()
		result = checkStringsFromStdIN(stdin, asciiSet, *printConciseFlag)
	}

	// output to file if given, else to STDOUT
	if len(*fileOut) > 0 {
		outputResults(*fileOut, result)
	} else {
		fmt.Print(result)
	}

}
