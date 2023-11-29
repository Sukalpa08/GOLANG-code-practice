package main

import (
	"fmt"
)

func Assigncookies() {
	var g [4]int = [4]int{9, 8, 7, 6}
	var s [3]int = [3]int{5, 6, 4}

	var (
		si, gi int
	)

	for gi < len(g) && si < len(s) {
		fmt.Println("cookies")
		if s[si] >= g[gi] {
			fmt.Println("Cookies not assigned")
			gi++
			si++
		} else {
			fmt.Println("Assigning cookies")
			gi++
		}
	}

}

func Foursum() {
	var A [2]int = [2]int{9, 8}
	var B [3]int = [3]int{7, 6, 5}
	var C [4]int = [4]int{4, 3, 2, 1}
	var D [0]int = [0]int{}
	var (
		record = make(map[int]int)
		res    int
	)

	for _, i := range A {
		for _, j := range B {
			fmt.Println("The sum of 4 numbers")
			record[i+j]++
		}
	}

	for _, i := range C {
		for _, j := range D {
			if s, ok := record[-i-j]; ok && s > 0 {
				fmt.Println("The sum of all numbers", res)
			}
		}
	}
}

func ContainDuplicate() {
	var nums [2]int = [2]int{1, 5}
	record := make(map[int]int)

	for _, num := range nums {
		fmt.Println("Duplicate")

		if _, ok := record[num]; !ok {
			fmt.Println("The number contains duplicate")
		} else {
			fmt.Println("The number does not contains duplicate")

		}
	}
}

func NoOfBoomerangs() {
	var points [][]int = make([][]int, 3)
	var n = len(points)

	points[0] = make([]int, 2)
	points[1] = make([]int, 2)
	points[2] = make([]int, 4)

	for i := 0; i < n; i++ {
		rr := make(map[int]int)
		fmt.Println("The record value")
		for j := 0; j < n; j++ {
			fmt.Println("The array value")
			if j != i {
				fmt.Println("Count the number of boomerangs")

			}
		}
		for _, j := range rr {
			fmt.Println("Boomerangs", j)
		}
	}
}

func KthLargest() {
	var nums [5]int = [5]int{0, 1, 2, 3, 4}
	var k int
	n := len(nums)

	for i := (n - 1) / 2; i >= 0; i-- {
		fmt.Println("The largest element of k")
	}

	for i := n - 1; i >= n-k; i-- {
		fmt.Println("The smallest element of k")

	}
}

func main() {
	Assigncookies()
	Foursum()
	ContainDuplicate()
	NoOfBoomerangs()
	KthLargest()
}
