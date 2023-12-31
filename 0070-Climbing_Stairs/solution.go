package main

func anotherClimbStairs(n int) int {
	ways := make([]int, n+1)
	ways[0] = 1

	for i := 1; i <= n; i++ {
		ways[i] += ways[i-1]
		if i >= 2 {
			ways[i] += ways[i-2]
		}
	}
	return ways[n]
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
