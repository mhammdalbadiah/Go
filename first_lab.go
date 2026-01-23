package main

import (
	"bufio" // To Enter FUll line
	"fmt"   // To Print
	"os"
)

func main() {

	congrats := "Enter Your Name"      // Declaer VAR
	Input := bufio.NewReader(os.Stdin) //  To Declear The VAR For The Input
	fmt.Println(congrats)
	name, _ := Input.ReadString('\n') // Input Just a Name you can Change it

	fmt.Println(name, "!")

}
