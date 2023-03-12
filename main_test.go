package main

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		input  []int
		shift  int
		expect int
	}{
		{[]int{4, 11, 3, 2, 59, 80}, 2, 85},
		{[]int{2, 11, 4, 20, 59, 80}, 2, 85},
		{[]int{2, 11, 4, 20, 59, 80}, 3, 85},
		{[]int{2, 11, 4, 20, 59, 2, 80}, 3, 142},
	}

	for i, v := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			result := MaxProfit(v.input, v.shift)

			if result != v.expect {
				t.Error("Expect: ", v.expect)
			}
		})
	}
	// 1. buy = 4, sell = 11, profit = 9
	// 2. buy = 2, sell = 80, profit = 78
	// expectaion = 85

	// 1. buy = 2, sell = 11, profit = 9
	// 2. buy = 4, sell = 80, profit = 76
	// 3. stop di share stok 2, karena profitnya lebih besar.
	// expectation = 85

	// 1. buy = 2, sell = 11, profit = 9
	// 2. buy = 4, sell = 59, profit = 55
	// 3. buy = 2, sell = 80, profit 78
	// expectation = 142
}

func TestShiftArray(t *testing.T) {
	testCases := []struct {
		input  []int
		shift  int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1, []int{4, 1, 2, 7, 5, 3, 8, 9, 6}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2, []int{7, 4, 1, 8, 5, 2, 9, 6, 3}},
	}

	for i, v := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			results := ShiftArray(v.input, v.shift)

			for j := range results {
				if results[j] != v.expect[j] {
					fmt.Println(v.expect)
					t.Error("Expect: ", v.expect)
				}
			}
		})
	}
}
