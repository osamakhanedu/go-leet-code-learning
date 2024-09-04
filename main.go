package main

import "fmt"

func containsDuplicate(nums []int) bool {

	var m = map[int]bool{}

	for _, value := range nums {
		fmt.Println(value)
		m[value] = true
	}

	return len(nums) != len(m)
}

func main() {

	nums := []int{1, 2, 3, 1}

	v := containsDuplicate(nums)

	fmt.Println(v)
}
