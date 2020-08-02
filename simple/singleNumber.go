package main

import "fmt"

func main() {
	nums := []int{9, 9, 7, 2, 7, 9, 7}
	fmt.Println(singleNumber(nums))

	nums2 := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(single2Number(nums2))
}

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		count := 0
		for _, n := range nums {
			//fmt.Println(1<<i, 1<<i&n)
			if 1<<i&n > 0 {
				count++
			}
		}
		if count%3 != 0 {
			res |= 1 << i
		}
	}
	return res
}

func single2Number(nums []int) []int {
	s := 0
	for _, n := range nums {
		s ^= n
	}
	l := s & (-s)
	fmt.Printf("%08b, %08b, %08b\n", s, -s, l)

	//fmt.Println(s,l)
	var z1, z2 int
	for _, n := range nums {
		fmt.Printf("%b, %b", l, n)
		if l&n == l {
			z1 ^= n
			fmt.Printf(" (l&n == l) z1=%b\n", z1)
		} else {
			z2 ^= n
			fmt.Printf(" (l&n != l) z2=%b\n", z2)
		}
	}
	return []int{z1, z2}
}
