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



// ============== Terminal ==================
// cd ~/ "File name "      : to enter a File 
// cd ..    : to go back 
// go run "FILE NAME "  : to run a Program 
// ls   : to print the Files 
// pwd  : to print your Location 
// clear : to clear The Terminal
// Ctrl + C : To Stop The Program
//
// ====== GIT =======
// git add .
// git commit -m "update"
// git push
// ==========================================





// There is a Different Between " Println , Printf " 
// 
// Println : Simple and jump a line at the End 
//
// Printf  : More Complex , That u can take the Full Control ,Doesn't Jump a line 



