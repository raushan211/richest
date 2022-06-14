package main

import "fmt"

func main() {
	input := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	fmt.Println(maximumWealth(input))
}

func maximumWealth(accounts [][]int) int {
	answer := 0
	for i := 0; i < len(accounts); i++ {
		element := accounts[i]
		sum := 0
		for j := 0; j < len(element); j++ {
			sum = sum + element[j]

		}
		if answer < sum {
			answer = sum
		}
	}
	return answer

}
