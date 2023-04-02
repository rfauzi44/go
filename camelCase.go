// Camel Case is a naming style common in many programming languages. In Java, method and variable names typically start with a lowercase letter, with all subsequent words starting with a capital letter (example: startThread). Names of classes follow the same pattern, except that they start with a capital letter (example: BlueCar).
// Your task is to write a program that creates or splits Camel Case variable, method, and class names.
// Input Format
// Each line of the input file will begin with an operation (S or C) followed by a semi-colon followed by M, C, or V followed by a semi-colon followed by the words you'll need to operate on.
// The operation will either be S (split) or C (combine)
// M indicates method, C indicates class, and V indicates variable
// In the case of a split operation, the words will be a camel case method, class or variable name that you need to split into a space-delimited list of words starting with a lowercase letter.
// In the case of a combine operation, the words will be a space-delimited list of words starting with lowercase letters that you need to combine into the appropriate camel case String. Methods should end with an empty set of parentheses to differentiate them from variable names.
// Output Format
// For each input line, your program should print either the space-delimited list of words (in the case of a split operation) or the appropriate camel case string (in the case of a combine operation).
// Sample Input

// S;M;plasticCup()

// C;V;mobile phone

// C;C;coffee machine

// S;C;LargeSoftwareBook

// C;M;white sheet of paper

// S;V;pictureFrame

// Sample Output

// plastic cup

// mobilePhone

// CoffeeMachine

// large software book

// whiteSheetOfPaper()

// picture frame

package main

import (
	"fmt"
	"strings"
	"unicode"
)

// variable = lowerCCase
// method = lowerCCase()
// class = CapitalizeCase

// split = loop, make sure, everyhting must be lowercase, base on camel, if there is camel add space

// combine = looping, make sure there is no space, if ther is space, move and camel
// C=67
// S=83
// V=86
// M=77

func main() {

	var C byte = 67
	var S byte = 83
	// var V byte = 86
	var M byte = 77
	var end byte = 41

	var inputLetter string = "S;V;epsonPrinter"
	letter := inputLetter[4:len(inputLetter)]

	if inputLetter[0] == C {

		words := strings.Split(letter, " ")

		for i := 1; i < len(words); i++ {

			words[i] = strings.ToUpper(string(words[i][0])) + words[i][1:]

		}

		wordsJoin := strings.Join(words, " ")

		spaceRemove := strings.ReplaceAll(wordsJoin, " ", "")

		if inputLetter[2] == M {
			parenthese := "()"
			result := spaceRemove + parenthese
			fmt.Println(result)

		} else if inputLetter[2] == C {
			if len(spaceRemove) > 0 {
				first := []rune(spaceRemove)[0]
				if unicode.IsLower(first) {
					first = unicode.ToUpper(first)
					output := string(first) + spaceRemove[1:]
					fmt.Println(output) // Output: Hello world
					return
				}
			}
			fmt.Println(spaceRemove)

		} else {
			fmt.Println(spaceRemove)
		}

	}

	if inputLetter[0] == S {

		var sanitize string

		if letter[len(letter)-1] == end {

			sanitize = letter[0:(len(letter) - 2)]

		} else {
			sanitize = letter
		}

		var output strings.Builder
		for i, r := range sanitize {
			if i > 0 && unicode.IsUpper(r) {
				output.WriteRune(' ')
			}
			output.WriteRune(r)
		}
		result := output.String()
		lowerResult := strings.ToLower(result)
		fmt.Println(lowerResult)

	}

}
