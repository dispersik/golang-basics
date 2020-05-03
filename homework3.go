package main

import (
	"fmt"
)

func average(arr [6]int) float64 {
	sum := 0
	for _, element := range arr {
		sum += element
	}
	return float64(sum) / 6
}

// max returns the longest word from
// the slice of strings
func max(s []string) string {
	maxIndex := 0
	maxLen := len(s[maxIndex])

	for i := 1; i < len(s); i++ {
		if len(s[i]) > maxLen {
			maxIndex = i
			maxLen = len(s[i])
		}
	}
	return s[maxIndex]
}

// reverse returns the copy of the original
// slice in reverse order. Type of slice elements is
// int64.
func reverse(arr []int64) []int64 {
	l := len(arr)
	reversed_arr := make([]int64, l, cap(arr))
	for i := 0; i < l; i++ {
		reversed_arr[i] = arr[l-1-i]
	}
	return reversed_arr
}

func main() {

	a := [6]int{6, 0, 0, 0, 0, 0}
	fmt.Println("avrage() func tests\n")
	fmt.Println("In: ", a)
	fmt.Println("Out: ", average(a), "\n")
	a = [6]int{0, 0, 0, 0, 0, 0}
	fmt.Println("In: ", a)
	fmt.Println("Out: ", average(a), "\n")

	fmt.Println("max() func tests\n")
	str := []string{"one", "zero", "five", "three"}
	fmt.Printf("In: %s\n", str)
	fmt.Printf("Out: %s\n\n", max(str))

	fmt.Println("reverse() func tests\n")
	arr := []int64{0,1,2,3,4}
	fmt.Printf("In: %v len = %v cap = %v\n", arr, len(arr), cap(arr))
	rev_arr:= reverse(arr)
	fmt.Printf("Out: %v len = %v cap = %v\n", rev_arr, len(rev_arr), cap(rev_arr))
	arr = make([]int64,2,18)
	arr = append(arr,1)
	fmt.Printf("In: %v len = %v cap = %v\n", arr, len(arr), cap(arr))
	rev_arr = reverse(arr)
	fmt.Printf("Out: %v len = %v cap = %v\n", rev_arr, len(rev_arr), cap(rev_arr))
}
