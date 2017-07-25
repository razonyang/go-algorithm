package replace_words

import (
	"strings"
)

func replaceWords(dict []string, sentence string) string {
	words := strings.Split(sentence, " ")

	for i, word := range words {
		for _, v := range dict {
			if len(word) >= len(v) && word[:len(v)] == v {
				words[i] = v
			}
		}
	}

	return strings.Join(words, " ")
}
