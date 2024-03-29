package main

import (
	"fmt"
)

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var count int32 = 0

	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			sum := ar[i] + ar[j]
			if sum%k == 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	result := divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2})
	fmt.Println(result)
}
