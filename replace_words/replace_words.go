package replace_words

import (
	"strings"
)

func replaceWords(dict []string, sentence string) string {
	words := strings.Split(sentence, " ")

	for i, word := range words {
		for _, v := range dict {
			if len(word) >= len(v) && word[:len(v)] == v {
				// in order to find the shortest root, doesn't break on here
				words[i] = v
			}
		}
	}

	return strings.Join(words, " ")
}
