// This Program is Very Useful for College Students !

package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

var courses int

// rff : Read From File
func rff() error {

	file, err := os.Open("Data.bin")
	if err != nil {
		return err
	}
	defer file.Close()

	dec := gob.NewDecoder(file)

	var hours []float32
	var grades []float32

	if err := dec.Decode(&hours); err != nil {
		return err
	}

	if err := dec.Decode(&grades); err != nil {
		return err
	}

	fmt.Println("Data From The File:")
	for i := 0; i < len(hours); i++ {
		fmt.Printf("Course %d -> Hours: %.1f | Grade: %.1f\n",
			i+1, hours[i], grades[i])
	}

	return nil
}

// stv : Save to File
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

// Cal_GPA : Calculate GPA
func Cal_GPA() float32 {

	cou := 0
	var GPA float32
	var SumG float32 = 0.0
	var SumH float32 = 0.0
	fmt.Println("Please Enter The Number Of Courses you have : ")
	fmt.Scan(&cou)
	courses = cou

	hours := make([]float32, courses)
	Grades := make([]float32, courses)

	for i := 0; i < courses; i++ {

		fmt.Println("Please Enter The Credits For ", i+1, ":")
		fmt.Scan(&hours[i])
		fmt.Println("Please Enter Your Grade :")
		fmt.Scan(&Grades[i])

	}

	GPA = 0
	for i := 0; i < courses; i++ {
		GPA += Grades[i] * hours[i]
	}
	GPA = GPA / SumH

	if err := stv(hours, Grades); err != nil {
		fmt.Println("Error storing data:", err)
	} else {
		fmt.Println("Data stored successfully")
	}

	return GPA

}

func main() {

	choice := 0
	fmt.Println("======= WELCOME ======")
	fmt.Println(" 1)  To Calculate Your GPA ")
	fmt.Println(" 2)  To Print Your Schedule")
	fmt.Println(" 3)  To Add Your Schedule ")
	fmt.Println(" 4)  To Exit 				")
	fmt.Println("======================")
	fmt.Printf(":")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		fmt.Println(Cal_GPA())
	case 2:

	case 3:

	case 4:

	}

}
