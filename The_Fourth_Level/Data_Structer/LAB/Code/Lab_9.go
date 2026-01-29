// This Program is Very Useful for College Students !

package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func stv(hours []float32, Grades []float32) error {

	file, err := os.OpenFile("Data.bin", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	enc := gob.NewEncoder(file)

	if enc := enc.Encode(hours); enc != nil {
		return err
	}
	if enc := enc.Encode(Grades); enc != nil {
		return err
	}

	return nil

}

func Cal_GPA() float32 {

	courses := 0
	var GPA float32
	var SumG float32 = 0.0
	var SumH float32 = 0.0
	fmt.Println("Please Enter The Number Of Courses you have : ")
	fmt.Scan(&courses)
	hours := make([]float32, courses)
	Grades := make([]float32, courses)

	for i := 0; i < courses; i++ {

		fmt.Println("Please Enter The Credits For ", i+1, ":")
		fmt.Scan(&hours[i])
		fmt.Println("Please Enter Your Grade :")
		fmt.Scan(&Grades[i])

	}

	for i := 0; i < courses; i++ {
		SumG = SumG + Grades[i]
		SumH = SumH + hours[i]
	}

	GPA = (SumG + SumH) / SumG

	return GPA

}

func main() {

	choice := 0
	fmt.Println("======= WELCOME ======")
	fmt.Println(" 1)  To Calculate Your GPA ")
	fmt.Println(" 2)  To Print Your Schedule")
	fmt.Println(" 3)  To Add Your Schedule ")
	fmt.Println(" 4)  To Exit 				")
	fmt.Scan(&choice)

	switch choice {

	case 1:

	case 2:

	case 3:

	case 4:

	}

}
