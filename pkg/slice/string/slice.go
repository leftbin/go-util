package string

import "strings"

func ContainsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ToMultiLineString(s []string) string {
	var b strings.Builder
	for i, e := range s {
		b.WriteString(e)
		if i == len(s)-1 {
			break
		}
		b.WriteString("\n")
	}
	return b.String()
}

func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
