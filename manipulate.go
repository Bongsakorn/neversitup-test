package main

// Manipulate function
func Manipulate(text string) []string {
	results := map[string]bool{}
	retResp := []string{}

	permutation(text, 0, len(text)-1, &results)

	for k, v := range results {
		if v {
			retResp = append(retResp, k)
		}

	}

	return retResp
}

func permutation(text string, strIdx, endIdx int, results *map[string]bool) {
	if strIdx == endIdx {
		(*results)[text] = true
	} else {
		for i := strIdx; i <= endIdx; i++ {
			text = swap(text, strIdx, i)
			permutation(text, strIdx+1, endIdx, results)
			text = swap(text, strIdx, i)
		}
	}
}

func swap(str string, i int, j int) string {
	runes := []rune(str)
	runes[i], runes[j] = runes[j], runes[i]
	return string(runes)
}

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}
