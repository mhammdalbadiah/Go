package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name  string
	ID    int
	Age   int
	Major string
}

func enterUsers(b []User) {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(b); i++ {
		fmt.Println("Enter User ID :")
		fmt.Scan(&b[i].ID)
		reader.ReadString('\n')

		fmt.Println("Enter User Name :")
		b[i].Name, _ = reader.ReadString('\n')
		b[i].Name = strings.TrimSpace(b[i].Name)

		fmt.Println("Enter User Age :")
		fmt.Scan(&b[i].Age)
		reader.ReadString('\n')

		fmt.Println("Enter User Major : ")
		b[i].Major, _ = reader.ReadString('\n')
		b[i].Major = strings.TrimSpace(b[i].Major)
	}

	fmt.Println("Nice , Do you Want To Print a Report ? (1)")
	report := 0
	fmt.Scan(&report)

	if report == 1 {

		fmt.Println("=========================================")
		for i := 0; i < len(b); i++ {
			fmt.Println("=========================================")
			fmt.Println("  User Number :", i+1)
			fmt.Println("  User ID :", b[i].ID)
			fmt.Println("  User Name :", b[i].Name)
			fmt.Println("  User Age :", b[i].Age)
			fmt.Println("  User Major :", b[i].Major)
			fmt.Println("=========================================")
		}
	}
}

func main() {
	fmt.Print("Please Enter Number Of Users : ")
	num := 0
	fmt.Scan(&num)

	users := make([]User, num)
	enterUsers(users)
}
