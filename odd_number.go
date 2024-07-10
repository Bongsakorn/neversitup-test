package main

// FindOddNumber finds the odd number in the given slice of integers.
func FindOddNumber(text []int) int {
	var oddNumber int

	for _, num := range text {
		oddNumber ^= num
	}
	// fmt.Println("number occur in odd time is ", oddNumber)
	return oddNumber
}
