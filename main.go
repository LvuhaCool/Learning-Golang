package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower = 2
	var userHeight float64
	var userWeight float64
	fmt.Println("___ Калькулятор индекса массы тела ___")
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userWeight)
	BMI := userWeight / math.Pow(userHeight, BMIPower)
	fmt.Printf("Ваш индекс массы тела: %v", BMI)
}
