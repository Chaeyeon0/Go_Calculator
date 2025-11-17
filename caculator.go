package gocalculator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Add(input string) (int, error) {
	if strings.TrimSpace(input) == "" {
		return 0, nil
	}

	delimiters := ",|:" // 기본 구분자
	numbers := input

	// 커스텀 구분자 형식 검사
	if strings.HasPrefix(input, "//") {
		re := regexp.MustCompile(`^//(.)\n(.*)`)
		matches := re.FindStringSubmatch(input)
		if len(matches) == 3 {
			delimiters = regexp.QuoteMeta(matches[1])
			numbers = matches[2]
		} else {
			return 0, errors.New("잘못된 구분자 형식입니다")
		}
	}

	re := regexp.MustCompile(delimiters)
	tokens := re.Split(numbers, -1)

	sum := 0
	for _, t := range tokens {
		num, err := strconv.Atoi(strings.TrimSpace(t))
		if err != nil {
			return 0, err
		}
		sum += num
	}

	return sum, nil
}
