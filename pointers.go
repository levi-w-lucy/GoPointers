package main

import "fmt"

func main() {
	age := 32
	var adultYears int

	fmt.Println("Age:", age)
	setAdultYears(&adultYears, age)
	fmt.Println("Adult years:", age)
}

func setAdultYears(adultYears *int, age int) {
	*adultYears = age - 18
}
