package gocalculator

import "testing"

func TestAdd_Basic(t *testing.T) {
	result, err := Add("1,2,3")
	if err != nil {
		t.Fatalf("에러: %v", err)
	}
	if result != 6 {
		t.Errorf("결과가 %d입니다. 기대값은 6이에요.", result)
	}
}

func TestAdd_CustomDelimiter(t *testing.T) {
	result, err := Add("//;\n1;2;3")
	if err != nil {
		t.Fatalf("에러 발생: %v", err)
	}
	if result != 6 {
		t.Errorf("결과가 %d입니다. 기대값은 6이에요.", result)
	}
}
