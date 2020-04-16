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

func max(s []string) string {
	// function that returns the longest word from
	// the slice of strings (the first if there are more
	// than one).
	// Input -> ("one", "two", "three")
	// Output -> ("three")
	// Input -> ("one", "two")
	// Output -> ("one")
	k := 0
	k_len := len(s[k])

	for i := 1; i < len(s); i++ {
		if len(s[i]) > k_len {
			k = i
			k_len = len(s[i])
		}
	}
	return s[k]
}

func reverse(arr []int64) []int64 {
	// function that returns the copy of the original
	// slice in reverse order. The type of elements is
	// int64.
	// Input -> (1, 2, 5, 15)
	// Output -> (15, 5, 2, 1)
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
	fmt.Println("In: ", arr, "len = ", len(arr), "cap = ", cap(arr))
	rev_arr:= reverse(arr)
	fmt.Println("Out: ", rev_arr, "len = ", len(rev_arr), "cap = ", cap(rev_arr), "\n")
	arr = make([]int64,2,18)
	arr = append(arr,1)
	fmt.Println("In: ", arr, "len = ", len(arr), "cap = ", cap(arr))
	rev_arr = reverse(arr)
	fmt.Println("Out: ", rev_arr, "len = ", len(rev_arr), "cap = ", cap(rev_arr))
}
