package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	var s []int
	fmt.Println("len", len(s))
	if s == nil {
		fmt.Println("nil slice")
	}
	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:4]
	fmt.Printf("s3 = %#v\n", s3)

	s3 = append(s3, 100)
	fmt.Printf("s3 (append) = %#v\n", s3)
	fmt.Printf("s2 (append) = %#v\n", s2) // s2 is changed as well
	fmt.Printf("s3: len = %d, cap = %d\n", len(s3), cap(s3))
	fmt.Printf("s2: len = %d, cap = %d\n", len(s2), cap(s2))

	// var s4 []int
	s4 := make([]int, 0, 1_000)
	for i := 0; i < 1_000; i++ {
		s4 = appendInt(s4, i)
	}
	fmt.Printf("s4: len = %d, cap = %d\n", len(s4), cap(s4))
	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D", "E"}))

	vs := []float64{2, 1, 3}
	fmt.Println(median(vs))
	vs = []float64{2, 1, 3, 4}
	fmt.Println(median(vs))
	fmt.Println(vs)

	fmt.Println(median(nil))
	fmt.Println(reflect.TypeOf(2))
	fmt.Println(reflect.TypeOf(2.0))
}

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}

	nums := make([]float64, len(values))
	copy(nums, values) // copy to not affect original

	sort.Float64s(nums)
	i := len(nums) / 2
	// if len(nums)&1 == 1 {
	if len(nums)%2 == 1 {
		return nums[i], nil
	}
	return (nums[i-1] + nums[i]) / 2, nil
}

func concat(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s

	// return append(s1, s2...)
}

func appendInt(s []int, v int) []int {
	i := len(s)

	if len(s) < cap(s) {
		s = s[:len(s)+1]
	} else {
		fmt.Printf("Reallocate: %d -> %d\n", len(s), len(s)*2+1)
		s2 := make([]int, len(s)*2+1)
		copy(s2, s)
		s = s2[:len(s)+1]
	}

	s[i] = v
	return s
}
