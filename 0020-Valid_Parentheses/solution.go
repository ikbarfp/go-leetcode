package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	pairs := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	var stack []rune
	for _, v := range s {
		if _, ok := pairs[v]; ok {
			stack = append(stack, v)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != v {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func anotherIsValid(s string) bool {
	guideline := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	flag := false

	for i := 0; i < len(s); i++ {
		currStr := fmt.Sprintf("%c", rune(s[i]))

		fmt.Println("currStr :", currStr)
		fmt.Println("pair :", guideline[currStr])

		if val, found := guideline[currStr]; found {

			for j := i + 1; j < len(s); j++ {
				fmt.Println()
				fmt.Println("s[j] :", string(s[j]))
				fmt.Println("val :", val)

				if _, currFound := guideline[string(s[j])]; currFound {
					if string(s[j]) == val {
						flag = true
					} else {
						flag = false
					}
				}
				continue
			}
		}
	}

	return flag
}
