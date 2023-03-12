package main

import "fmt"

func main() {
	fmt.Println("====== Test 1 =======")
	fmt.Println(ShiftArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1))

	fmt.Println()

	fmt.Println("====== Test 2 =======")
	fmt.Println(MaxProfit([]int{2, 11, 4, 20, 59, 2, 80}, 3))
}
