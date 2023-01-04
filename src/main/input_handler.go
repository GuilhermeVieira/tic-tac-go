package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func readValidPlayerInput() int64 {
	var input string

	for {
		input = scanPlayerInput()
		if isInputValid(input) {
			break
		}
		fmt.Println("Warning: Must be between 1 and 9. Please try again!")
	}

	return parseInt(input)
}

func scanPlayerInput() string {
	var input string

	for {
		_, err := fmt.Scanln(&input)
		if err == nil {
			return input
		}
		log.Println(err)
		fmt.Println("Warning: There was an error while reading the input. Please try again!")
	}
}

func isInputValid(input string) bool {
	integerBetweenOneAndNine := regexp.MustCompile(`\b[1-9]\b`)
	return integerBetweenOneAndNine.MatchString(input)
}

func parseInt(s string) int64 {
	res, err := strconv.ParseInt(s, 0, 32)

	if err != nil {
		log.Fatalln("There was a fatal error while reading player input.")
	}

	return res
}
