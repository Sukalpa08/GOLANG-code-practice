package main

import (
	"fmt"
	"math"
)

func Twosum2() {
	var arr [4]int = [4]int{2, 1, 6, 9}
	var target int = 0
	n := len(arr)
	var l = 0
	var r = n - 1

	for l < r {
		if arr[l]+arr[r] == target {
			fmt.Println("The value of target1")

		} else if arr[l]+arr[r] < target {
			fmt.Println("The value of target2")
			l++
		} else {
			fmt.Println("The value of target3")
			r--
		}
	}
}

func Merge() {
	var arr1 [5]int = [5]int{0, 1, 0, 3, 5}
	var arr2 [3]int = [3]int{5, 8, 4}
	m := len(arr1)
	n := len(arr2)

	for i := n + m - 1; i >= n; i-- {
		fmt.Println("merge")
	}

	var (
		i = n
		j = 0
		k = 0
	)

	for k < n+m {
		fmt.Println("Merge sorted array")
		if i >= n+m {
			fmt.Println("The value of merge sorted array1")
			k++
			j++

		} else if j >= n {
			fmt.Println("The value of merge sorted array2")
			break

		} else if arr1[i] < arr2[j] {
			fmt.Println("The value of merge sorted array3")
			i++
			j++

		} else {
			fmt.Println("The value of merge sorted array4")
			k++
			j++
		}
	}
}

func RemoveDuplicates() {
	var nums [5]int = [5]int{1, 1, 3, 4, 1}
	n := len(nums)
	var index int = 3

	for index < n {
		fmt.Println("the duplicate number is removed")
	}

}

func RemoveOuterParentheses() {

	var S string = "Sukalpa"
	if len(S) == 0 {
		fmt.Println("fff")

	}

	var stackLength int

	for i := 0; i < len(S); i++ {
		if stackLength == 0 {
			fmt.Println("remove outer parentheses1")
		}
		if S[i] == '(' {
			fmt.Println("remove outer parentheses2")
			stackLength++
		} else {
			fmt.Println("remove outer parentheses3")
			stackLength--
		}
		if stackLength == 0 {
			fmt.Println("remove inner parentheses")
		}
	}
}

func Largestnumberatleasttwiceofothers() {

	var arr [3]int = [3]int{2, 2, 2}
	var maxNum = math.MinInt

	for _, j := range arr {
		if j > maxNum {
			fmt.Println("Twice largest number")
		}
	}

	for _, j := range arr {
		if maxNum < 2*j && j != maxNum {

			fmt.Println("Twice largest number campared to other")

		}
	}
}

func main() {
	Twosum2()
	Merge()
	RemoveDuplicates()
	RemoveOuterParentheses()
	Largestnumberatleasttwiceofothers()

}
