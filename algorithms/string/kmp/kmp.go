package kmp

func buildPatternTable(word string) []int {
	patternTable := make([]int, len(word))
	prefixIndex, suffixIndex := 0, 1
	for suffixIndex < len(word) {
		if word[prefixIndex] == word[suffixIndex] {
			patternTable[suffixIndex] = prefixIndex + 1
			suffixIndex++
			prefixIndex++
		} else if prefixIndex == 0 {
			suffixIndex++
		} else {
			prefixIndex = patternTable[prefixIndex-1]
		}
	}
	return patternTable
}

func KMP(text, word string) int {
	if len(word) == 0 {
		return 0
	}

	textIndex, wordIndex := 0, 0
	patternTable := buildPatternTable(word)

	for textIndex < len(text) {
		if text[textIndex] == word[wordIndex] {
			if wordIndex == len(word)-1 {
				return textIndex - len(word) + 1
			}
			wordIndex++
			textIndex++
		} else if wordIndex > 0 {
			wordIndex = patternTable[wordIndex-1]
		} else {
			wordIndex = 0
			textIndex++
		}
	}
	return -1
}
