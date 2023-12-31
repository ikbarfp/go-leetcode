package main

import "strings"

// "Big O thing matters" solution
func lengthOfLastWord(s string) int {
	counter := 0
	for i := len(s) - 1; i >= 0; i-- {
		if counter == 0 && s[i] == ' ' {
			continue
		} else {
			if s[i] == ' ' {
				break
			}

			counter++
		}
	}
	return counter
}

// Naive solution
func anotherLengthOfLastWord(s string) int {
	arr := strings.Split(strings.TrimSpace(s), " ")

	return len(arr[len(arr)-1])
}
