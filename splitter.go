package gocalculator

import (
	"regexp"
	"strings"
)

func SplitNumbers(input string) []string {
	if strings.TrimSpace(input) == "" {
		return []string{}
	}

	delimiters := ",|:"
	numbers := input

	if strings.HasPrefix(input, "//") {
		re := regexp.MustCompile(`^//(.)\n(.*)`)
		matches := re.FindStringSubmatch(input)
		if len(matches) == 3 {
			delimiters = regexp.QuoteMeta(matches[1])
			numbers = matches[2]
		}
	}

	re := regexp.MustCompile(delimiters)
	return re.Split(numbers, -1)
}
