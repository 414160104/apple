package main

import (
	"fmt"
	"sort"
)

func main() {
	timeRange1 := []int{6, 7}
	timeRanges := [][]int{
		{1, 5},
		{8, 10},
	}

	if checkConflict(timeRange1, timeRanges) {
		fmt.Println("Conflict detected!")
	} else {
		fmt.Println("No conflict detected!")
	}

	segments := [][]int{{1, 8}, {7, 10}, {14, 18}}
	checkConflicts(segments)
	fmt.Println(IsFoundInRecord("", nil))

}

func checkConflict(timeRange1 []int, timeRanges [][]int) bool {
	for _, timeRange2 := range timeRanges {
		if rangeOverlap(timeRange1, timeRange2) {
			return true
		}
	}
	return false
}

func rangeOverlap(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a[0] > b[1]) || (b[0] > a[1]) {
		return false
	}
	return true
}

func checkConflicts(segments [][]int) bool {
	sort.Slice(segments, func(i, j int) bool {
		return segments[i][0] < segments[j][0]
	})

	for i := 1; i < len(segments); i++ {
		if overlap(segments[i-1], segments[i]) {
			fmt.Println("conflict")
			return true
		}
	}

	return false
}

func overlap(a, b []int) bool {
	return a[1] > b[0] && b[1] > a[0]
}

func IsFoundInRecord(searchString string, stringArray []string) bool {
	strArray := []string{"apple", "banana", "cherry", "date", "grape"}

	srchString := "che1rry"

	found := false
	for _, str := range strArray {
		if str == srchString {
			found = true
			break
		}
	}

	return found
}
