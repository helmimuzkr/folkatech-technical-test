package main

// First Apporach
// O(N2) Time Complexity
func ShiftArray(array []int, i int) []int {
	n := len(array)
	// If i or the shift number greater than the length of array,
	// shifting array equal to i mod n.
	shift := i % n

	for shift != 0 {
		tempArr := make([]int, n)
		// Middle element.
		tempArr[4] = array[4]

		// top-right rotation.
		for index := 0; index < n/2; index++ {
			swap := (index + 1) % n
			if index > 1 {
				swap = index*2 + (index - 1)
			}
			if index == 3 {
				index += 2
			}
			tempArr[swap] = array[index]
		}

		// bottom-left rotation.
		for index := n - 1; index > 0; index-- {
			swap := (index - 1) % n
			if index < 7 {
				if index == 6 {
					swap -= 2
				} else {
					swap -= 4
				}
			}
			if index == 5 {
				index -= 2
			}
			tempArr[swap] = array[index]
			if swap == 0 {
				break
			}
		}

		// Update the array.
		array = tempArr
		shift--
	}
	return array
}

// Second Apporach
// O(N) Time Complexity
// func ShiftArray(array []int, i int) []int {
// 	n := len(array)
// 	// if i or the shift number greater than the length of array,
// 	// shifting array equal to i mod n.
// 	shift := i % n
// 	if shift == 0 {
// 		return array
// 	}

// 	for shift != 0 {
// 		tempArr := make([]int, n)
// 		// To Right.
// 		tempArr[1], tempArr[2], tempArr[5], tempArr[8] = array[0], array[1], array[2], array[5]
// 		// Middle.
// 		tempArr[4] = array[4]
// 		// To Left.
// 		tempArr[7], tempArr[6], tempArr[3], tempArr[0] = array[8], array[7], array[6], array[3]

// 		// Update the array.
// 		array = tempArr
// 		shift--

// 		// for j := range array {
// 		// 	if j%3 == 0 && j != 0 {
// 		// 		fmt.Println()
// 		// 	}
// 		// 	fmt.Printf("%d ", array[j])
// 		// }
// 		// fmt.Println()
// 	}

// 	return array
// }

// https://jamboard.google.com/d/1F6J6f2pWTJ-GMrVGuQUB6CAGnvvnn8mPOLNSAzuG_F8/viewer?f=0
