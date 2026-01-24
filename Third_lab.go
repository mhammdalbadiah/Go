package main
  
import (
	"fmt"
)

func main() {

	// GO have only " For loop "

	// First Type :
	for i := 0; i < 10; i++ {
		fmt.Println("I love Go")
	}
	// Second Type :
	counter := 0
	for {
		fmt.Println("I love Go")
		if counter > 10 {
			break
		}
		counter++

	}

}
