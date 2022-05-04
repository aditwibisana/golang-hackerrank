package main

import "fmt"

func server(arr[] int, size int) {
	for i := 0; i < size; i++ {
		fmt.Println(" ",arr[i])
	}

}
func main() {
	var arr[] int
	fmt.Scan(&arr)
	var size int = len(arr)
	server(arr, size)
}