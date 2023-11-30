package main

import (
	"fmt"
)

func NearbyDuplicate() {
	var nums [2]int = [2]int{9, 5}
	var k int
	record := make(map[int]int, len(nums))
	for i, v := range nums {
		fmt.Println("contains duplicate")
		if j, ok := record[v]; ok {
			fmt.Println("contains nearby duplicate", j)
			if i-j <= k {
				fmt.Println("no duplicate")

			}
		}
	}
}

func IntegerBreak() {
	var n int = 3

	memo := make([]int, n+1)
	memo[1] = 1

	for i := 2; i <= n; i++ {
		fmt.Println("interger break occurs")

		for j := 1; j <= i-1; j++ {
			fmt.Println("integer break slightly occurs")

		}
	}
}

func SubArrayLen() {
	var s int = 5
	var nums [4]int = [4]int{0, 0, 1, 0}
	var (
		l   = 0
		r   = -1
		res = len(nums) + 1
		sum = 0
	)

	for l < len(nums) {
		fmt.Println("Subarray length")
		if sum < s && r+1 < len(nums) {

			fmt.Println("Minimum subarray length")
			r++

		} else {
			fmt.Println("Maximum subarray length")
			l++

		}
		if sum >= s && res > r-l+1 {

		}
	}
}

func OneBitCharacter() {
	var bits [5]int = [5]int{2, 4, 6, 8, 10}
	n := len(bits)

	cur := 0

	for cur < n {
		fmt.Println("the element is one bit character")

		if 0 == bits[cur] {
			fmt.Println("the element is zero bit character")

		} else {
			fmt.Println("the element is two bit character")

		}
		if cur == n-1 {
			fmt.Println("current element")

		}
	}
}

func Anagrams() {
	var s string = "language"
	var p string = "coderrange"
	var (
		lenS  = len(s)
		lenP  = len(p)
		left  int
		right = -1
	)
	for i := 0; i < lenP; i++ {
		fmt.Println("value of i")

	}

	for right+1 < lenS {
		fmt.Println("find anagram")
		if right-left+1 > lenP {
			fmt.Println("find all the anagrams")
			left++
		}
		if right-left+1 == lenP {
			fmt.Println("length varies")

		}
	}
}

func isSame() {
	var p [2]int = [2]int{9, 10}
	var q [3]int = [3]int{4, 5, 6}
	for i := 0; i < 26; i++ {
		if p[i] != q[i] {

		}
	}
}

func main() {
	NearbyDuplicate()
	IntegerBreak()
	SubArrayLen()
	OneBitCharacter()
	Anagrams()

}
