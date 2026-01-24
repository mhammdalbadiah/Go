package main

import (
	"fmt"
)

func main() {

	num := 0
	fmt.Println("Enter The Size Of the Array : ")
	fmt.Scan(&num)

	arr := make([]int8, num)
	for i := 0; i < num; i++ {

		fmt.Println("Array Num :", i+1)
		fmt.Scan(&arr[i])

	}

}
