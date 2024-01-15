package main

func migratoryBirds(arr []int32) int32 {
	frequency := make(map[int32]int)

	for _, num := range arr {
		frequency[num]++
	}

	var smallestFrequentElement int32 = 0
	mostFrequent := 0

	for num, freq := range frequency {
		if freq > mostFrequent {
			mostFrequent = freq
			smallestFrequentElement = num
		} else if freq == mostFrequent && num < smallestFrequentElement {
			smallestFrequentElement = num
		}
	}

	return smallestFrequentElement
}

func main() {
	result := migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4})
	println(result)
}
