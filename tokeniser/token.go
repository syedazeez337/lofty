package tokeniser

import (
	"fmt"
	"regexp"
	"strconv"
)

type Tokens struct {
	nums       int
	whitespace string
	operators  string
}

func currentToken(position int, word string) byte {
	if position >= len(word) {
		return 0
	}

	return word[position]
}

func DisplayMsg(word string) []string {
	setOfinput := make([]string, len(word))

	for i, elem := range word {
		setOfinput[i] = string(elem)
	}

	return setOfinput
}

func SpecialFunc(word string) []int {
	out := filter(word)
	return out
}

func filter(word string) []int {
	var res []int

	// Regular expression to match integers
	re := regexp.MustCompile(`\d+`)

	// Find all matches
	matches := re.FindAllString(word, -1)

	for _, elem := range matches {

		elemInt, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Println("Error", err)
		}
		res = append(res, elemInt)
	}
	return res
}
