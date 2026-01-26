package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tasks []string
var dates []string

// Add Task
func Add_Task(reader *bufio.Reader) {
	fmt.Print("Enter The Task: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	fmt.Print("Enter The Date: ")
	date, _ := reader.ReadString('\n')
	date = strings.TrimSpace(date)

	if task == "" || date == "" {
		fmt.Println("Task and Date cannot be empty")
		return
	}

	tasks = append(tasks, task)
	dates = append(dates, date)

	fmt.Println("Task Added")
}

// Show Tasks
func Show_Tasks() {
	if len(tasks) == 0 {
		fmt.Println("There is No Tasks!")
		return
	}

	fmt.Println("==== Tasks ====")
	for i := 0; i < len(tasks); i++ {
		fmt.Println(i+1, ")", tasks[i], "On", dates[i])
	}
}

// Delete Task
func Delete_Task(reader *bufio.Reader) {
	if len(tasks) == 0 {
		fmt.Println("No Tasks to Delete!")
		return
	}

	Show_Tasks()
	fmt.Print("Enter Task Number: ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n, err := strconv.Atoi(line)
	if err != nil || n < 1 || n > len(tasks) {
		fmt.Println("Invalid Number")
		return
	}

	index := n - 1
	tasks = append(tasks[:index], tasks[index+1:]...)
	dates = append(dates[:index], dates[index+1:]...)

	fmt.Println("Task Deleted")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== TO DO LIST =====")
		fmt.Println("1) Add Task")
		fmt.Println("2) Show Tasks")
		fmt.Println("3) Delete Task")
		fmt.Println("4) Exit")
		fmt.Print("Choose: ")

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		choice, _ := strconv.Atoi(line)

		switch choice {
		case 1:
			Add_Task(reader)
		case 2:
			Show_Tasks()
		case 3:
			Delete_Task(reader)
		case 4:
			fmt.Println("Bye")
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}
}
