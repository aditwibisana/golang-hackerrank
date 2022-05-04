package main

import (
	"fmt"
)

// Compare the Triplets
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/compare-the-triplets/problem

func main() {
	var a [3]int
	var b [3]int
	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Scan(&b[0], &b[1], &b[2])
	var bob, alice int
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			alice++
		} else if a[i] < b[i] {
			bob++
		}
	}
	fmt.Println(alice, bob)
	//fmt.Println(strings.Trim(fmt.Sprint(result), "[]")) // menghilangkan kurung siku
}
