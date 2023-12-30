package main

import "strings"

func romanToInt(s string) int {
	counter := 0
	arrOfNum := strings.Split(s, "")

	for i := 0; i < len(arrOfNum); i++ {
		switch arrOfNum[i] {
		case "I":
			counter += 1
		case "V":
			if i != 0 && arrOfNum[i-1] == "I" {
				counter -= 1
				counter += 4
			} else {
				counter += 5
			}
		case "X":
			if i != 0 && arrOfNum[i-1] == "I" {
				counter -= 1
				counter += 9
			} else {
				counter += 10
			}
		case "L":
			if i != 0 && arrOfNum[i-1] == "X" {
				counter -= 10
				counter += 40
			} else {
				counter += 50
			}
		case "C":
			if i != 0 && arrOfNum[i-1] == "X" {
				counter -= 10
				counter += 90
			} else {
				counter += 100
			}
		case "D":
			if i != 0 && arrOfNum[i-1] == "C" {
				counter -= 100
				counter += 400
			} else {
				counter += 500
			}
		case "M":
			if i != 0 && arrOfNum[i-1] == "C" {
				counter -= 100
				counter += 900
			} else {
				counter += 1000
			}
		}
	}

	return counter
}
