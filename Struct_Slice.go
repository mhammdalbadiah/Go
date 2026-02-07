package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct { // This Struct For The Slice
	Id   int
	Name string
	Dept string
	GPA  float64
}

type Class struct { // This one is for The big Struct , Which Include "Student" Struct
	Class_Num int
	Floor_Num int
	Students  []Student
}

// ===== Helper Functions For Input =====
func ReadLine(reader *bufio.Reader, msg string) string {
	fmt.Print(msg)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func ReadInt(reader *bufio.Reader, msg string) int {
	for {
		text := ReadLine(reader, msg)
		val, err := strconv.Atoi(text)
		if err == nil {
			return val
		}
		fmt.Println("Sorry , INVALID NUMBER ! Try Again.")
	}
}

func ReadFloat(reader *bufio.Reader, msg string) float64 {
	for {
		text := ReadLine(reader, msg)
		val, err := strconv.ParseFloat(text, 64)
		if err == nil {
			return val
		}
		fmt.Println("Sorry , INVALID NUMBER ! Try Again.")
	}
}

// ===== Students Functions =====

// To Enter The Student Info Using Slice With Struct !
func Enter_Student_Info(reader *bufio.Reader, a []Student) {
	for i := 0; i < len(a); i++ {
		fmt.Println("======        Student ", i+1, "        ======")

		a[i].Id = ReadInt(reader, "Please Enter Student ID : ")
		a[i].Name = ReadLine(reader, "Please Enter Student Name : ")
		a[i].Dept = ReadLine(reader, "Please Enter Student Department : ")
		a[i].GPA = ReadFloat(reader, "Please Enter Student GPA : ")

		fmt.Println("==========================================")
	}
}

// To Print Students Info If it Exist
func Print_Student_Info(a []Student) {
	if len(a) <= 0 {
		fmt.Println("Sorry , There is NO Students Info !")
		return
	}

	for i := 0; i < len(a); i++ {
		fmt.Println("=========== Student Number :", i+1, "============")
		fmt.Println("ID :", a[i].Id)
		fmt.Println("Name :", a[i].Name)
		fmt.Println("Department :", a[i].Dept)
		fmt.Println("GPA :", a[i].GPA)
		fmt.Println("==============================================")
	}
}

// ===== Class Functions =====

func Add_Class_Info(reader *bufio.Reader, classes *[]Class, students []Student) {
	if len(students) <= 0 {
		fmt.Println("Sorry , You Must Add Students First !")
		return
	}

	c := Class{}
	fmt.Println("========== Add Class Info ==========")
	c.Class_Num = ReadInt(reader, "Please Enter Class Number : ")
	c.Floor_Num = ReadInt(reader, "Please Enter Floor Number : ")

	// Copy Students List Into This Class
	c.Students = make([]Student, len(students))
	copy(c.Students, students)

	*classes = append(*classes, c)
	fmt.Println("Class Added Successfully !")
	fmt.Println("====================================")
}

func Print_Class_Info(classes []Class) {
	if len(classes) <= 0 {
		fmt.Println("Sorry , There is NO Classes Info !")
		return
	}

	for i := 0; i < len(classes); i++ {
		fmt.Println("=============== Class Number :", i+1, "===============")
		fmt.Println("Class_Num :", classes[i].Class_Num)
		fmt.Println("Floor_Num :", classes[i].Floor_Num)
		fmt.Println("Students Count :", len(classes[i].Students))
		fmt.Println("----------------------------------------------")

		// Print Students Inside This Class
		if len(classes[i].Students) > 0 {
			for j := 0; j < len(classes[i].Students); j++ {
				s := classes[i].Students[j]
				fmt.Println("  -> Student", j+1)
				fmt.Println("     ID :", s.Id)
				fmt.Println("     Name :", s.Name)
				fmt.Println("     Dept :", s.Dept)
				fmt.Println("     GPA :", s.GPA)
				fmt.Println("     ------------------------")
			}
		} else {
			fmt.Println("No Students In This Class !")
		}

		fmt.Println("====================================================")
	}
}

// ===== Main =====

func main() {
	reader := bufio.NewReader(os.Stdin)

	var students []Student
	var classes []Class

	for {
		fmt.Println("======= WelCome To Our Program ======== ")
		fmt.Println("1) Add Student Info")
		fmt.Println("2) Print Student Info")
		fmt.Println("3) Add Class Info")
		fmt.Println("4) Print Class Info")
		fmt.Println("5) EXIT")
		fmt.Println("======================================= ")

		choice := ReadInt(reader, ":")

		switch choice {

		case 1:
			sNum := ReadInt(reader, "Enter The Number Of Student : ")
			if sNum <= 0 {
				fmt.Println("Sorry , INVALID NUMBER !")
				break
			}

			students = make([]Student, sNum)
			Enter_Student_Info(reader, students)

		case 2:
			fmt.Println("Here a Full Report About The Student")
			Print_Student_Info(students)

		case 3:
			Add_Class_Info(reader, &classes, students)

		case 4:
			Print_Class_Info(classes)

		case 5:
			fmt.Println("Good Bye !")
			return

		default:
			fmt.Println("INVALID NUMBER !")
		}

		fmt.Println()
	}
}
