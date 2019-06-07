package main

import (
	"fmt"
)

func main() {
	// START OMIT
	n1, n2 := 0.1, 0.2
	n3 := 0.3
	n4 := n3 - n1 - n2
	if n4 == 0 {
		fmt.Println("n4 equals zero")
	} else {
		fmt.Println("n4 not equals zero")
	}
	// END OMIT
}
