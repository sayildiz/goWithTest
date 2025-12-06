package iteration

import "strings"

const repeatCount = 5

func RepeatConcat(s string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += s
	}
	return repeated
}

func RepeatBuilder(s string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}
