package main
 
import (
	"fmt"
)

func main() {

	id := 0 // Or var id int8

	fmt.Println("Enter Your ID : ")
	fmt.Scan(&id)

	if id > 234 {

		fmt.Println(id, "Bigger Than", 234)
	} else if 234 > id {
		fmt.Println(234, "Bigger Than ", id)

	} else {
		fmt.Println(id, "And", 234, " They are Equal !")
	}

}
