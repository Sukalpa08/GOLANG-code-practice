package main

import (
	"fmt"
	"math"
)

func Twosum() {
	var arr [4]int = [4]int{1, 2, 2, 5}
	var target int = 0
	for i, j := range arr {
		fmt.Println("two elements", j)
		for a := i + 1; a < len(arr); a++ {
			fmt.Println("sum of values")
			if arr[a]+j == target {
				fmt.Println("target value")

			}
		}

	}

}

func IsPalindrome() {
	var x int
	var (
		y int
		z = x
	)

	for y > 0 {

	}
	fmt.Println(y == z)
}

func Reverseinteger() {
	var x int = 3
	if 0 == x {

	}
	for x < 0 {
		fmt.Println("x gives -ve value")
	}

	for x > 0 {
		fmt.Println("x gives +ve value")
	}
}

func MaxSubArray() {

	var arr [3]int = [3]int{3, 4, 2}
	var max = math.MinInt32
	var sum int = 3
	for _, j := range arr {
		fmt.Println("The subarray value", j)

		if sum > max {
			fmt.Println("Maximum subarray value")

		}
		if sum < 0 {
			fmt.Println("Minimum subarray value")
		}
	}

}

func Movezero() {

	var arr [4]int = [4]int{1, 2, 2, 3}
	var k int = 0
	for i := range arr {
		fmt.Println("array value")
		if arr[i] != 0 {
			fmt.Println("move zero forward")

			if k != i {
				fmt.Println("move zero")

			}
			fmt.Println("move zero", arr[i])
			k++
		}
	}

}

func main() {
	//Twosum()
	//IsPalindrome()
	//Reverseinteger()
	//MaxSubArray()
	//Movezero()
}
