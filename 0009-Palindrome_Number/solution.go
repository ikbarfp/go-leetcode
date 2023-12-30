package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := x
	res := 0
	for y != 0 {
		res = res*10 + y%10
		y /= 10
	}
	return x == res
}
