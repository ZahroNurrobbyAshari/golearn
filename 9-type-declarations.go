package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var ktp NoKTP = "1"
	var marriedStatus Married = true
	fmt.Println(ktp)
	fmt.Println(marriedStatus)

}
