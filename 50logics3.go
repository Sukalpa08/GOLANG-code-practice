package main

import (
	"fmt"
)

func Singlenumber() {

	var nums [3]int = [3]int{1, 5, 9}
	res := 0
	for _, num := range nums {
		res ^= num
		fmt.Println("Singlenumber")
	}

}

func Maxprofit() {

	var prices [3]int = [3]int{1, 3, 5}

	for i := 0; i < len(prices); i++ {

		fmt.Println("Profit")

		for j := 0; j < i; j++ {

			fmt.Println("Max profit")

			if prices[i]-prices[j] > 0 {

			}
		}
	}
}

func MaxArea() {
	var height [4]int = [4]int{1, 3, 5, 2}

	for i := 0; i < len(height); i++ {
		fmt.Println("height of a container")

		for j := i + 1; j < len(height); j++ {
			fmt.Println("height of a water")

		}
	}
}

func RemoveElement() {
	var nums [4]int = [4]int{0, 1, 1, 3}
	var val int
	x := 0
	for j := 0; j < len(nums); j++ {
		fmt.Println("remove the element from array")
		if nums[j] != val {
			fmt.Println("value of the element")
			if x != j {
				fmt.Println("campare that element")

			}
		}
	}
}

func LastWord() {
	var s string = "levels"
	var n = len(s)
	var cur = n - 1

	for cur >= 0 {
		if n-1 == cur && s[cur] == 32 {
			fmt.Println("length of last word")
			cur--
			n--

		}
		if s[cur] != 32 {
			fmt.Println("length of the word")
			cur--

		} else {
			fmt.Println("nothing")

		}
	}
}

func main() {
	Singlenumber()
	Maxprofit()
	MaxArea()
	RemoveElement()
	LastWord()

}
