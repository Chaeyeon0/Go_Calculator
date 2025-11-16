package gocalculator

import (
	"strconv"
	"strings"
)

func Add(input string) (int, error) {
	if strings.TrimSpace(input) == "" {
		return 0, nil
	}

	// 쉼표 또는 콜론으로 분리
	parts := strings.FieldsFunc(input, func(r rune) bool {
		return r == ',' || r == ':'
	})

	sum := 0
	for _, p := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return 0, err
		}
		sum += num
	}
	return sum, nil
}
