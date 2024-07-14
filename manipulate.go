package main

import (
	"sort"
	"strings"
)

var manipulateArr []string
var checkDuplicate map[string]struct{}

func Manipulate(text []string) []string {

	checkDuplicate = make(map[string]struct{})

	manipulate(text, 0)
	// slices.Sort(manipulateArr) // for go version 1.21+
	sort.Strings(sort.StringSlice(manipulateArr))

	defer func() {
		// set zero
		manipulateArr = []string{}
		checkDuplicate = map[string]struct{}{}
	}()

	return manipulateArr
}

func manipulate(arr []string, index int) {
	length := len(arr)
	if index == length-1 {
		str := strings.Join(arr, "")
		if _, exist := checkDuplicate[str]; !exist {
			manipulateArr = append(manipulateArr, str)
			checkDuplicate[str] = struct{}{}
		}
	} else {
		for i := index; i < length; i++ {
			arr[index], arr[i] = arr[i], arr[index]
			manipulate(arr, index+1)
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
}
