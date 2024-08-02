//go:build ignore

package generator

import (
	"regexp"
	"strings"
	"unicode"
)

//###########################################################//

func toTitleCase(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	r := regexp.MustCompile(re)
	words := strings.Fields(s)

	for i, word := range words {
		if len(word) > 0 {
			word = r.ReplaceAllString(word, "")
			if len(word) > 0 {
				words[i] = string(unicode.ToUpper(rune(word[0]))) + word[1:]
			}

		}
	}

	return strings.Join(words, "")
}

func head(pack string) string {
	return "package " + pack + "\n\n" + hText + "\n//######################################//\n\n"
}
