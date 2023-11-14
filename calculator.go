// begin      Эта версия калькулятора пока что работает только с арбаскими цифрами.
package main

import "fmt"

func main() {
	fmt.Println("Привет, это калькулятор.\n Форма выглядит следующим образом: x + y ") //Пример ввода для пользователя
	var x, y int
	var c string           // Инициализация переменных
	fmt.Scanln(&x, &c, &y) // Ввод пользователем данных и присвоение этих значений переменным
	switch c {             //Проверка операций, проверка же числовых данных будет в отдельных функциях.
	case "+":
		sum(x, y)
		break
	case "-":
		raz(x, y)
		break
	case "*":
		umn(x, y)
		break
	case "/":
		del(x, y)
		break
	default:
		fmt.Println("Error") // Нет нужной операции, вывод ошибки
	}

}
func sum(x, y int) { // Если вместо числа пользователь напишет буквы значение останется дэфолтным и будет 0,не выполнив условия программа выведет ошибку,так же с отрицательными числами
	if x <= 10 && x >= 1 && y <= 10 && y >= 1 {
		fmt.Println(x + y)
	} else {
		fmt.Println("Error")
	}
}
func umn(x, y int) {
	if x <= 10 && x >= 1 && y <= 10 && y >= 1 {
		fmt.Println(x * y)
	} else {
		fmt.Println("Error")
	}
}
func raz(x, y int) {
	if x <= 10 && x >= 1 && y <= 10 && y >= 1 {
		fmt.Println(x - y)
	} else {
		fmt.Println("Error")
	}
}
func del(x, y int) {
	if x <= 10 && x >= 1 && y <= 10 && y >= 1 {
		fmt.Println(x / y)
	} else {
		fmt.Println("Error")
	}
}
