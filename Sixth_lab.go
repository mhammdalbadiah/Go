package main

import (
	"fmt"
)

// Global VAR
var balance int64 = 1000

// Deposit
func Deposit() {
	var Dep int64
	fmt.Println("Please Enter The Amount you want to Deposit : ")
	fmt.Scan(&Dep)
	balance += Dep
	fmt.Println("DEPOSIT DONE !")

}

// Withdraw
func Withdraw() {
	var wit int64
	fmt.Println("Please Enter The Amount you want To Withdraw :")
	fmt.Scan(&wit)
	if wit <= balance {
		balance -= wit
		fmt.Println("WITHDRAW DONE !")
	} else {
		fmt.Println("INVALID!")
	}

}

// Check Balance
func Check() {
	fmt.Println("YOUR BALANCE IS :", balance)
}

func main() {
	choice := 1

	for {

		fmt.Println("================================")
		fmt.Println("	WLLCOME TO ALINMA BANK")
		fmt.Println(" 1. Deposit              ")
		fmt.Println(" 2. Withdraw			  ")
		fmt.Println(" 3. Check Balance.       ")
		fmt.Println(" 4. Exit 				  ")
		fmt.Println("================================")
		fmt.Scan(&choice)
		if choice == 4 {
			break
		}

		switch choice {

		case 1:
			Deposit()
		case 2:
			Withdraw()
		case 3:
			Check()
		}
	}

}
