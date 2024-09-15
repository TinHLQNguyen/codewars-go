package spinwords

import "strings"

/// NOTE:
//Write a function that takes in a string of one or more words, and returns the same string,
//but with all words that have five or more letters reversed (Just like the name of this Kata).
//Strings passed in will consist of only letters and spaces. Spaces will be included only when more than one word is present.

var REVERSETHRESHOLD = 5

func SpinWords(input string) string {
	words := strings.Split(input, " ")
	for i, word := range words {
		if len(word) >= REVERSETHRESHOLD {
			words[i] = reverseWord(word)
		}
	}
	return strings.Join(words, " ")
}

func reverseWord(input string) string {
	// FIXME: this method is known to not working on combining characters (Examples: Vietnamese vowel)

	// Template runes to the max possible length
	// b/c len(input) returns length of bytes, and each rune (loosely means char) can be encoded by >1 bytes
	runes := make([]rune, len(input))
	// use runecount to actually figure out the number of runes in the input
	runecount := 0
	// range(input) actually iterate by rune's code point, not per bytes
	for _, r := range input {
		runes[runecount] = r
		runecount++
	}
	// trim the original template to the actual runecount
	runes = runes[0:runecount]

	// Now we can work on runes like normal
	for i := 0; i < runecount/2; i++ {
		runes[i], runes[runecount-1-i] = runes[runecount-1-i], runes[i]
	}

	// After swapping, now convert back to string
	return string(runes)
}
