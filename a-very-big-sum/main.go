package main

import (
	"fmt"
)

// A Very Big Sum
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/a-very-big-sum/problem

func main() {

	var angka, temp, res int
	fmt.Scan(&angka)
	for i := 0; i < angka; i++ {
		fmt.Scan(&temp)
		res += temp
	}
	fmt.Println(res)
}
