package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userWeight float64
	fmt.Println("___ Калькулятор индекса массы тела ___")
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userWeight)
	IMT := userWeight / math.Pow(userHeight, IMTPower)
	fmt.Printf("Ваш индекс массы тела: %v", IMT)
}
