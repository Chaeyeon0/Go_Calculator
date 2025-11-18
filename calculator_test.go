package gocalculator

import "testing"

func TestSum(t *testing.T) {
	result := Sum([]int{1, 2, 3})
	if result != 6 {
		t.Errorf("결과가 %d 입니다. 기대값은 6이에요.", result)
	}
}
