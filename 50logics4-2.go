package main

import (
	"fmt"
)

func UniquePaths() {
	var m int = 2
	var n int = 4

	for i := 0; i < m; i++ {
		fmt.Println("make unique pathsm")

	}

	for i := 1; i < n; i++ {
		fmt.Println("make unique pathsn")

	}

	for i := 1; i < m; i++ {
		fmt.Println("Unique paths of m")
		for j := 1; j < n; j++ {
			fmt.Println("Unique paths of n")

		}
	}
}

func Triangle() {
	var triangle [][]int = make([][]int, 3)
	m := len(triangle)
	dp := make([][]int, m)

	triangle[0] = make([]int, 2)
	triangle[1] = make([]int, 2)
	triangle[2] = make([]int, 3)

	for i := 1; i < m; i++ {
		fmt.Println("The value of i")
		for j, _ := range triangle[i] {
			fmt.Println("The value of j")
			if j >= len(dp[i-1]) {
				fmt.Println("triangle")

			} else {
				if j-1 >= 0 && dp[i-1][j-1] < dp[i-1][j] {
					fmt.Println("minimum length of triangle")

				} else {
					fmt.Println("nothing")

				}
			}
		}
	}
}

func SortColors() {
	var nums [5]int = [5]int{1, 2, 3, 4, 5}
	var two = len(nums)

	for i := 0; i < two; i++ {
		fmt.Println("sort colors")
		if 1 == nums[i] {
			fmt.Println("Quick way to sort colors")

		} else if 2 == nums[i] {
			fmt.Println("lazy way to sort colors")

		} else {
			fmt.Println("nothing")

		}
	}
}

func PlusOne() {
	var digits [3]int = [3]int{7, 7, 8}
	var carry = 1
	var index = len(digits) - 1

	for index >= 0 {
		fmt.Println("adding a digit")
		index--
	}

	if carry > 0 {
		fmt.Println("no adding digit")

	}
}

func HouseRob() {
	var nums [4]int = [4]int{4, 5, 11, 5}
	n := len(nums)

	for i := 0; i < n; i++ {
		fmt.Println("More chances of house robbering")

	}

	for i := n - 2; i >= 0; i-- {
		fmt.Println("Chances are less")
		for j := i; j < n; j++ {
			fmt.Println("Chances are normal")

			if j+2 < n {
				fmt.Println("no chances")

			} else {
				fmt.Println("nothing")

			}
		}
	}
}

func main() {
	UniquePaths()
	Triangle()
	SortColors()
	PlusOne()
	HouseRob()
}
