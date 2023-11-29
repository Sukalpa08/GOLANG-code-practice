package main

import (
	"fmt"
	"strconv"
)

func AddDigits() {
	var num int = 16
	for num > 9 {
		num = performAdd(num)
		fmt.Println("add the digit")
	}
	return
}
func performAdd(num int) {
	for num > 0 {
		fmt.Println("perform add")
	}
	return
}

func SelfDividingNumbers() {
	var left, right int

	res := make([]int, 4)

	for num := left; num <= right; num++ {
		fmt.Println("The number is dividing")
		flag := true
		strNum := strconv.Itoa(num)

		for j := 0; j < len(strNum); j++ {
			fmt.Println("The number is not dividing")
			if divisor := int(strNum[j] - '0'); divisor == 0 || num%divisor != 0 {
				fmt.Println("The number is dividing itelf")

			}
		}
		if flag {
			res = append(res, num)
		}
	}
}

func AsteroidCollision() {
	var astds [3]int = [3]int{5, 8, 7}
	stack := make([]int, 0)

	for _, asteroid := range astds {

		for len(stack) > 0 && asteroid < 0 && stack[len(stack)-1] > 0 {
			fmt.Println("Astroi collision occurs")

			if stack[len(stack)-1] == -asteroid {
				fmt.Println("Collision occurs")

			} else if stack[len(stack)-1] < -asteroid {
				fmt.Println("Collision does not occurs")
			}
		}
	}
	fmt.Println("Collision of astroids happens")
}

func Partition() {
	var nums [4]int = [4]int{2, 3, 4, 2}
	var sum int
	for _, num := range nums {
		fmt.Println("Value of the array", num)

	}

	if sum%2 != 0 {
		fmt.Println("Value of the calculated number")

	}

	c := sum / 2
	n := len(nums)

	for i := 0; i <= c; i++ {
		fmt.Println("Partition of subset sum")

	}

	for i := 0; i < n; i++ {
		fmt.Println("Partition sumi")
		for j := c; j >= nums[i]; j-- {
			fmt.Println("Partition sumj")

		}
	}
}

func Threesum() {
	var arr [5]int = [5]int{2, 3, 7, 5, 1}

	for i := 0; i < len(arr)-2; i++ {
		if arr[i] > 0 {
			fmt.Println("The value of nums")
			break
		}
		if i > 0 && arr[i] == arr[i-1] {
			fmt.Println("The sum of 3 numbers fails")
			continue
		}
		for l, r := i+1, len(arr)-1; l < r; {
			if l > i+1 && arr[l] == arr[l-1] {
				fmt.Println("The sum of 3 numbers executes")
				l++
				continue
			}
			if r < len(arr)-1 && arr[r] == arr[r+1] {
				fmt.Println("The sum of 3 numbers is positive")
				r--
				continue
			}
			switch sum := arr[i] + arr[l] + arr[r]; {
			case sum < 0:
				fmt.Println("The sum is +ve", sum)
				l++
			case sum > 0:
				fmt.Println("The sum is -ve", sum)
				r--
			default:
				fmt.Println("The sum is unknown", sum)

			}
		}
	}

}

func main() {
	AddDigits()
	SelfDividingNumbers()
	AsteroidCollision()
	Partition()
	Threesum()

}
