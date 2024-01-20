package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
	mapOfColors := make(map[int32]int32)
	var counter int32

	for _, color := range ar {
		mapOfColors[color]++
	}

	for _, count := range mapOfColors {
		counter += count / 2
	}

	return counter
}

func main() {
	fmt.Println(sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}))
}
