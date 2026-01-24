package main

import (
	"fmt"
)

// Golbal VAR
const N = 10

var Student [N]float64
var Max float64
var Min float64

// Input Student Marks
func Input_Marks() {
	fmt.Println("Enter Student Marks : ")
	for i := 0; i < N; i++ {
		fmt.Println("Student Number :", i+1)
		fmt.Scan(&Student[i])
	}
}

// OutPut Student Marks
func Output_Marks() {

	for i := 0; i < N; i++ {

		fmt.Println("Student Number :", i+1, "Has", Student[i])

	}

}

// AVG
func AVG_Marks() float64 {
	avg := 0.0
	for i := 0; i < N; i++ {

		avg += Student[i]

	}
	return avg / N

}

// Max
func MAX_Marks() float64 {
	max := Student[0]
	for i := 1; i < N; i++ {
		if Student[i] > max {
			max = Student[i]
		}
	}
	return max
}

// Min
func MIN_Marks() float64 {
	min := Student[0]
	for i := 1; i < N; i++ {
		if Student[i] < min {
			min = Student[i]
		}
	}
	return min
}

func main() {

	Input_Marks()
	Output_Marks()
	fmt.Println("AVG :", AVG_Marks())
	fmt.Println("MAX :", MAX_Marks())
	fmt.Println("MIN :", MIN_Marks())

}
