package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SearchInsert() {
	var nums [5]int = [5]int{0, 1, 5, 9, 11}
	var target int
	var (
		l int
		r = len(nums) - 1
	)

	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			fmt.Println("search number")
		}
		if target < nums[mid] {
			fmt.Println("search insert")
			r = mid - 1
		} else {
			fmt.Println("insrt number")
			l = mid + 1
		}
	}
	return
}

func RomanToInt() {
	var s string = "XVIC"

	digits := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100}
	sum := digits[s[len(s)-1]]

	for i := len(s) - 1; i > 0; i-- {
		cur := digits[s[i]]
		pre := digits[s[i-1]]
		if cur > pre {
			fmt.Println("s")

		} else {
			fmt.Println("z")

		}
	}
	fmt.Println("covert roman to integer", sum)
	return
}

func evalRPN() {
	var tokens [4]string = [4]string{"14", "6", "0", "1"}
	stack := make([]int, len(tokens))
	top := -1
	for i := 0; i < len(tokens); i++ {
		switch ch := tokens[i]; ch {
		case "+":
			stack[top-1] += stack[top]
			top--
		case "-":
			stack[top-1] -= stack[top]
			top--
		case "*":
			stack[top-1] *= stack[top]
			top--
		case "/":
			stack[top-1] /= stack[top]
			top--
		default:
			top++
		}
	}
	fmt.Println(stack[0])
	return
}

func CompareVersion() {
	var version1 string = "Version"
	var version2 string = "Comparison"
	var (
		v1 = strings.Split(version1, ".")
		v2 = strings.Split(version2, ".")
	)
	for i := 0; i < len(v1) || i < len(v2); i++ {
		var v1N, v2N int
		if i < len(v1) {
			fmt.Println("Numbers calculated")
			v1N, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			fmt.Println("Numbers compared")
			v2N, _ = strconv.Atoi(v2[i])
		}

		if v1N > v2N {
			fmt.Println("version 1")

		} else if v1N < v2N {
			fmt.Println("version 2")

		}
	}
	return
}

func Intersect() {
	var nums1 [4]int = [4]int{12, 43, 77, 100}
	var nums2 [5]int = [5]int{72, 55, 29, 110}
	record := make(map[int]int)
	res := make([]int, 0)
	for _, num := range nums1 {
		if _, ok := record[num]; !ok {
			fmt.Println("combining of 2 arrays")

		} else {
			fmt.Println("2 arrays incremented")
			record[num]++
		}
	}

	for _, num := range nums2 {
		if count, ok := record[num]; ok && count > 0 {
			fmt.Println("intersection of 2 arrays")
			res = append(res, num)
			record[num]--
		}
	}
	return
}

func main() {
	RomanToInt()
	SearchInsert()
	evalRPN()
	CompareVersion()
	Intersect()
}
