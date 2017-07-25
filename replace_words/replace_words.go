package replace_words

import (
	"strings"
)

func replaceWords(dict []string, sentence string) string {
	words := strings.Split(sentence, " ")

	for i := 0; i < len(words); i++ {
		for _, v := range dict {
			if len(words[i]) >= len(v) && words[i][:len(v)] == v {
				// in order to find the shortest root, doesn't break on here
				words[i] = v
			}
		}
	}

	return strings.Join(words, " ")
}
