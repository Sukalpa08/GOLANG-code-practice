package main

import (
	"fmt"
)

func NumSubArrayProductLessThanK() {
	var arr [4]int = [4]int{3, 5, 5, 6}
	var k int
	if k <= 1 {

	}
	var prod = 1
	var res = 0
	var left = 0

	for right, val := range arr {
		prod *= val
		for prod >= k {
			fmt.Println("Product less than k")
			left++
		}
		res += right - left + 1
	}
}

func Binarysearch() {
	var arr [3]int = [3]int{3, 7, 14}
	var target int
	var l int
	var r = len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] == target {

			fmt.Println("search the number")
		}
		if arr[mid] > target {
			fmt.Println("binary search")

		} else {
			fmt.Println("search the value")

		}

	}

}

func FindLengthOfLCIS1() {
	var arr [6]int = [6]int{1, 1, 5, 4, 4, 3}
	n := len(arr)
	if 0 == n || 1 == n {

	}

	var maxLen int
	var count = 1

	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			fmt.Println("subsequence value")
			count++
		} else {
			fmt.Println("subsequence  value")
		}

		if count > maxLen {
			fmt.Println("longest increasing subsequence")
		}
	}
}

func isSubsequence() {
	var s string = "code1"
	var t string = "code2"
	var si, ti int

	for si < len(s) && ti < len(t) {
		fmt.Println("code")
		if s[si] == t[ti] {
			fmt.Println("Sequence")
			si++
			ti++
		} else {
			fmt.Println("Subsequence")
			ti++
		}
	}
}

func ReverseString() {
	var s string = "code"
	var l int
	var r = len(s) - 1
	for r > l {
		fmt.Println("The string is reversed")
		r--
		l++
	}
}

func main() {

	NumSubArrayProductLessThanK()
	Binarysearch()
	FindLengthOfLCIS1()
	isSubsequence()
	ReverseString()
}
