package gocalculator

import (
	"reflect"
	"testing"
)

func TestSplitNumbers_DefaultDelimiters(t *testing.T) {
	result := SplitNumbers("1,2:3")
	expected := []string{"1", "2", "3"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("결과가 %v 입니다. 기대값은 %v", result, expected)
	}
}

func TestSplitNumbers_CustomDelimiter(t *testing.T) {
	result := SplitNumbers("//;\n1;2;3")
	expected := []string{"1", "2", "3"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("결과가 %v 입니다. 기대값은 %v", result, expected)
	}
}
