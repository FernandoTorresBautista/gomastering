/*
	Binary search implementation in Go
*/
package main

import "fmt"

func binarySearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}
	mid := int((lowIndex + highIndex) / 2)
	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

func iterBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int
	for startIndex < endIndex {
		mid = int((startIndex + endIndex) / 2)
		if array[mid] > target {
			endIndex = mid
		} else if array[mid] < target {
			startIndex = mid
		} else {
			return mid
		}
	}
	return -1
}

// extrac the code up from section 2> BinarySearch

type Set map[string]struct{}

func getSetValues(s Set) []string {
	var retVal []string
	for k, _ := range s {
		retVal = append(retVal, k)
	}
	return retVal
}

// extrac the code up from Section 2>Section2-6.go

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 6, 10, 15, 15, 18, 19, 30}

	find := binarySearch(slice, 10, 0, len(slice)-1)
	fmt.Println("Found element by recursive binary search: ", find)

	find = iterBinarySearch(slice, 10, 0, len(slice)-1)
	fmt.Println("Found element by recursive binary search: ", find)

	//
	s := make(Set)
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}
	//get and print items
	fmt.Println(getSetValues(s))
}
