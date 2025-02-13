package main

import (
	"fmt"
)

func main() {
	const maxCount = 15
	var number int

	fmt.Printf("Введите количество элементов массива (максимум new %d): ", maxCount)
	fmt.Scan(&number)

	if number > maxCount {
		fmt.Printf("Количество элементов не может превышать %d\n", maxCount)
		return
	}

	numArray := make([]int, number)

	fmt.Println("Введите элементы массива:")
	for i := 0; i < number; i++ {
		fmt.Scan(&numArray[i])
	}

	fmt.Println("Введенный массив:")
	for _, v := range numArray {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
