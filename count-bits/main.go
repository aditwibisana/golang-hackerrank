package main

import "fmt"

func counBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func main() {
	var n int
	count := 0
	fmt.Scan(&n)
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	fmt.Println(count)
	//n := 126
	//counBits(n)
	//fmt.Println(counBits(n))
}
