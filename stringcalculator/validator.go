package gocalculator

import (
	"errors"
	"strconv"
)

func ValidateNumbers(nums []string) ([]int, error) {
	result := []int{}
	for _, s := range nums {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.New("숫자가 아닌 문자가 포함되어 있습니다")
		}
		if n < 0 {
			return nil, errors.New("음수는 허용되지 않습니다")
		}
		result = append(result, n)
	}
	return result, nil
}
