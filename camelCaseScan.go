package main

import (
	"bufio"
	"fmt"
	"os"
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

	scanner := bufio.NewScanner(os.Stdin)

	// Read input from STDIN line by line
	for scanner.Scan() {
		input := scanner.Text()

		inputLetter := input // var inputLetter string = "C;M;mobile crap guo"
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

			} else {
				fmt.Println(spaceRemove)
			}

		}

		if inputLetter[0] == S {

			var output strings.Builder
			for i, r := range letter {
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

}

// C;V;can of coke
// S;M;sweatTea()
// S;V;epsonPrinter
// C;M;santa claus
// C;C;mirror

// canOfCoke
// sweat tea
// epson printer
// santaClaus()
// Mirror

// canOfCoke
// sweat tea()
// epson printer
// santaClaus()
// Mirror
