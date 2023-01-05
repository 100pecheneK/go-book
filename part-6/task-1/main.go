package main

import (
	"fmt"
	"sort"
)

func sortInt(a, b, c int) (p1, p2, p3 int) {
	nums := []int{a, b, c}
	sort.Ints(nums)
	p1, p2, p3 = nums[0], nums[1], nums[2]
	return
}
func sortInt2(a, b, c int) (int, int, int) {
	nums := []int{a, b, c}
	sort.Ints(nums)
	p1, p2, p3 := nums[0], nums[1], nums[2]
	return p1, p2, p3
}

func main() {
	p1, p2, p3 := sortInt(3, 2, 1)
	p12, p22, p32 := sortInt2(3, 2, 1)
	fmt.Println(p1, p2, p3)
	fmt.Println(p12, p22, p32)
}
