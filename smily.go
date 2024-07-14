package main

import (
	"regexp"
)

func CountSmilyFace(text []string) int {
	var count int

	r, err := regexp.Compile("[)D]")
	if err != nil {
		panic(err)
	}

	for _, v := range text {
		if r.MatchString(v) {
			count++
		}

	}
	return count
}
