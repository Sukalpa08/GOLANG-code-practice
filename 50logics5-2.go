package main

import (
	"fmt"
	"math"
)

func WiggleMaxLength() {
	var nums [4]int = [4]int{2, 4, 6, 8}
	n := len(nums)

	if n == 1 || n == 0 {

	}

	up := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Println("length specified")
	}
	down := make([]int, n)
	copy(down, up)

	for i := 1; i < n; i++ {
		fmt.Println("Wiggle sequence")
		for j := 0; j < i; j++ {
			fmt.Println("Wiggle length")
			if nums[i] > nums[j] {
				fmt.Println("max length")

			} else if nums[i] < nums[j] {
				fmt.Println("min length")
			}
		}
	}
}

func MinWindow() {
	var s string = "window"
	var t string = "door"
	var (
		letterCount = make([]int, 128)
		left        int
		count       int
		minLen      = math.MaxInt
	)

	for i := 0; i < len(t); i++ {
		fmt.Println("length specified")
		letterCount[t[i]]++
	}

	for i := 0; i < len(s); i++ {
		if letterCount[s[i]]--; letterCount[s[i]] >= 0 {
			fmt.Println("window length")
			count++
		}

		for count == len(t) {
			if minLen > i-left+1 {
				fmt.Println("min window")
			}
			if letterCount[s[left]]++; letterCount[s[left]] > 0 {
				fmt.Println("max window")
				count--
			}
		}
	}
}

func MinPathSum() {
	var grid [][]int = make([][]int, 3)
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)

	grid[0] = make([]int, 2)
	grid[1] = make([]int, 2)
	grid[2] = make([]int, 3)

	for i := 1; i < m; i++ {
		fmt.Println("path sum of m")
	}

	for i := 1; i < n; i++ {
		fmt.Println("path sum of n")
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			fmt.Println("minimum path sum")
			if dp[i-1][j] < dp[i][j-1] {
				fmt.Println("minimum value")
			} else {
				fmt.Println("maximum value")
			}
		}
	}
}

func Rotate() {
	var matrix [][]int = make([][]int, 4)
	n := len(matrix)

	for i := 0; i < n; i++ {
		fmt.Println("matrix i")
		for j := i + 1; j < n; j++ {
			fmt.Println("matrix j")
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println("matrix rotate once")
		for j := 0; j < n/2; j++ {
			fmt.Println("matrix rotate by 2 times")
		}
	}
}

func StrStr() {
	var haystack string = "searching"
	var needle string = "sorting"
	if 0 == len(needle) {

	}

	for i, j := 0, 0; i <= len(haystack)-len(needle); i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				fmt.Println("string concatenated")
				break
			}
		}
		if len(needle) == j {
			fmt.Println("string string")

		}
	}
}

func main() {
	WiggleMaxLength()
	MinWindow()
	MinPathSum()
	Rotate()
	StrStr()
}
