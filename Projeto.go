package main

import (
	"fmt"
	"math"
)

func main() {
	//Exer2()
	//Exer3()
	//Exer4()
	//Exer5()
	//Exer6(2)
}

func Exer2() {
	fmt.Println("Digite seu nome:")
	var input string

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("é um prazer te conhecer,", input, "!")
}

func Exer3() {
	fmt.Println("Digite um valor")
	var n1 int

	fmt.Scan(&n1)

	fmt.Println("Digite outro valor")
	var n2 int

	fmt.Scan(&n2)

	soma := n1 + n2

	println("A soma entre ", n1, "e", n2, "é igual a", soma)
}

// func Exer4() {
// 	fmt.Println("Digite algo")

// 	var input string
// 	fmt.Scan(&input)

// 	value := input
// 	fmt.Println(reflect.TypeOf(value))
// }

func Exer5() {
	fmt.Println("Digite um numero inteiro!")
	var value int
	fmt.Scan(&value)

	fmt.Println("O sucessor é ", Exer5a(value), "e o antecessor é", Exer5b(value))
}

func Exer5a(number int) int {

	n := number + 1
	return n

}
func Exer5b(number int) int {

	n := number - 1
	return n

}

func Exer6(number int) {
	fmt.Println("Digite um numero inteiro")

	var value int
	fmt.Scan(&value)

	println("O dobro de", value, "é", value*2)
	println("O triplo de", value, "é", value*3)

	value2 := float64(value)

	println("A raiz quadrada de", value, "é", math.Sqrt(value2))

}
