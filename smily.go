package main

import "regexp"

// CountSmilyFace counts the number of smily faces in the given slice of strings.
func CountSmilyFace(text []string) int {
	var validSmilyFace int
	smileyPattern := regexp.MustCompile(`^[:;][-~]?[D)]`)

	for _, face := range text {
		if smileyPattern.MatchString(face) {
			validSmilyFace++
		}
	}

	return validSmilyFace
}
