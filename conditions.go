// § Условия

package main

import "fmt"

func main() {
	a := 5
	b := 6
	if a < b {
		fmt.Println("a is less than b")
	} else if a > b {
		fmt.Println("a is more than b")
	} else {
		fmt.Println("a равно b")
	}
}

// есть еще конструкция switch case
