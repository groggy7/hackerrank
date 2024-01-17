package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	var itemsAnnaAte int32 = 0
	for i, item := range bill {
		if i == int(k) {
			continue
		}
		itemsAnnaAte += item
	}
	if b == itemsAnnaAte/2 {
		fmt.Print("Bon Appetit")
	} else {
		num := b - itemsAnnaAte/2
		fmt.Print(num)
	}
}

func main() {
	bonAppetit([]int32{3, 10, 2, 9}, 1, 7)
}
