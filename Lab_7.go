package main

import "fmt"

func main() {

	var num int
	change := 0

	fmt.Println("Welcome! Enter The Number of Elements:")
	fmt.Scan(&num)

	ele := make([]int, num)

	for i := 0; i < num; i++ {
		fmt.Println("Element Number:", i+1)
		fmt.Scan(&ele[i])
	}

	fmt.Println("If you want to Change the Size of the Elements, enter NEW size:")
	fmt.Scan(&change)

	if change > num {
		for i := 0; i < change-num; i++ {

			ele = append(ele, 0)

			fmt.Println("Element Number:", num+i+1)
			fmt.Scan(&ele[len(ele)-1])
		}
	} else {
		fmt.Println("INVALID INPUT !")
	}

	fmt.Println("Final Elements:", ele)
	fmt.Println("len:", len(ele), "cap:", cap(ele))
}
