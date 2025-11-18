package main

import (
	"bufio"
	"fmt"
	"gocalculator"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("덧셈할 문자열을 입력해 주세요:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	parts := gocalculator.SplitNumbers(text)
	validNums, err := gocalculator.ValidateNumbers(parts)
	if err != nil {
		fmt.Println("에러:", err)
		return
	}

	result := gocalculator.Sum(validNums)
	fmt.Println("결과 :", result)
}
