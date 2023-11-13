// begin
package main

import "fmt"

func main() {
	fmt.Println("Привет, это калькулятор.\n Форма выглядит следующим образом: x + y ")
	var x, y int
	var c string
	fmt.Println("Vvedite chislo")
	fmt.Scanln(&x, &c, &y)
	fmt.Println(sum(x, y))

}
func sum(x, y int) int {
	var sum int = 0
	if x <= 10 && x >= 1 && y <= 10 && y >= 1 {
		sum = x + y
	} else {
		fmt.Println("Sorry")
	}
	return sum
}
func umn(x, y int) int {
	var umn int = 0
	if x <= 10 && x >= 1 && y <= 10 && y >= 1 {
		umn = x * y
	} else {
		fmt.Println("Sorry")
	}
	return umn
}
func raz(x, y int) int {
	var raz int = 0
	if x <= 10 && x >= 1 && y <= 10 && y >= 1 {
		raz = x - y
	} else {
		fmt.Println("Sorry")
	}
	return raz
}
func del(x, y int) int {
	var del int = 0
	if x <= 10 && x >= 1 && y <= 10 && y >= 1 {
		del = x / y
	} else {
		fmt.Println("Sorry")
	}
	return del
}
