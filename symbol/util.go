package symbol

import "strings"

func JoinSentences(operator string, sentences ...*Sentence) string {
	var builder strings.Builder
	for i, sent := range sentences {
		builder.WriteString(sent.Description)

		if i != len(sentences)-1 {
			builder.WriteString(" ")
			builder.WriteString(operator)
		}
	}
	return builder.String()
}
