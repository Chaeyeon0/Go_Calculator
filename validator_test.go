package gocalculator

import "testing"

func TestValidateNumbers_Valid(t *testing.T) {
	nums, err := ValidateNumbers([]string{"1", "2", "3"})
	if err != nil {
		t.Fatalf("에러 발생: %v", err)
	}
	if len(nums) != 3 {
		t.Errorf("길이가 %d 입니다. 기대값은 3이에요.", len(nums))
	}
}

func TestValidateNumbers_Invalid(t *testing.T) {
	_, err := ValidateNumbers([]string{"1", "a"})
	if err == nil {
		t.Errorf("숫자가 아닌 입력에 에러가 발생해야 합니다.")
	}
}

func TestValidateNumbers_Negative(t *testing.T) {
	_, err := ValidateNumbers([]string{"1", "-2"})
	if err == nil {
		t.Errorf("음수 입력에 에러가 발생해야 합니다.")
	}
}
