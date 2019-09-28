package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Orange", "Plum", "Pomegranate", "Grape"}
	some := fruits[1:3:3] // length = 3-1 equals capacity = 3  ; default capacity = k-i / 5-1
	fmt.Println(cap(fruits))
	fmt.Println(cap(some))
	fmt.Println("o fruits : ", fruits)
	some = append(some, "Banana", "Mango")

	fmt.Println("n fruits : ", fruits)
	fmt.Println("some : ", some)
}
