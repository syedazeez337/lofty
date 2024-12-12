package tokeniser

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
