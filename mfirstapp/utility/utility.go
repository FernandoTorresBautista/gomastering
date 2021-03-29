package utility

import "fmt"

// exported function
func SaidHello() {
	fmt.Println("Hello World")
}

/* // is not exported function
func sayFer() {
	fmt.Println("Hello Fer")
} */

// is not exported function
func SayFer() {
	fmt.Println("Hello Fer")
}
