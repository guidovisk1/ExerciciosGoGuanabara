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
	// fmt.Println(Exerc7(20, 4))
	//Exer8()
	//Exerc9()
	//Exerc9a(12)
	//Exerc10()
	//Exerc11()
	//Exerc12()
	//Exerc13()
	Exerc14()
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

func Exerc7(n1 float64, n2 float64) float64 {

	media := Exerc7a(n1, n2)
	return media
}

func Exerc7a(num1, num2 float64) float64 { // Pegando a média
	return (num1 + num2) / 2
}

func Exer8() {
	var number float64

	fmt.Println("Digite uma distância em metros")
	fmt.Scan(&number)

	fmt.Printf("A conversão de %.3f metros é\n", number) // %.2f especifica duas casas decimais
	fmt.Printf("%.3f km /", number/1000)
	fmt.Printf("%.2f hm /", number/100)
	fmt.Printf("%.2f dam /", number/10)
	fmt.Printf("%.2f dm /", number*10)
	fmt.Printf("%.2f cm /", number*100)
	fmt.Printf("%.2f mm /", number*1000)

}

func Exerc9() {
	var number int
	fmt.Println("Digite um número")
	fmt.Scan(&number)
	fmt.Println(number, "x 1 =", number*1)
	fmt.Println(number, "x 2 =", number*2)
	fmt.Println(number, "x 3 =", number*3)
	fmt.Println(number, "x 4 =", number*4)
	fmt.Println(number, "x 5 =", number*5)
	fmt.Println(number, "x 6 =", number*6)
	fmt.Println(number, "x 7 =", number*7)
	fmt.Println(number, "x 8 =", number*8)
	fmt.Println(number, "x 9 =", number*9)
	fmt.Println(number, "x 10 =", number*10)
}
func Exerc9a(numero int) {
	var number int
	fmt.Println("Digite um número")
	fmt.Scan(&number)
	for i := 0; i < 11; i++ {
		fmt.Println(number, "*", i, "=", number*(i+1)-number)
	}
}

func Exerc10() {
	var numero float64

	fmt.Println("Quantos reais voce tem?")
	fmt.Scan(&numero)

	fmt.Println("Com R$", numero, "voce pode comparar USD", numero/5)
}

func Exerc11() {
	fmt.Println("Largura da parede")
	var n1 float64
	fmt.Scan(&n1)

	fmt.Println("Altura da parede")
	var n2 float64
	fmt.Scan(&n2)

	area := n1 * n2
	fmt.Println("A sua area tem", area, "m2")
	fmt.Println("Para pintar essa parede, você precisará de", area/2, "l de tinta")
}

func Exerc12() {
	var number float64
	fmt.Println("Digite um valor")
	fmt.Scan(&number)

	desconto := number * 0.05
	promocao := number - desconto

	fmt.Println("O produto que custava", number, "na promoção de 5% irá custar", promocao)
}

func Exerc13() {
	var number float64
	fmt.Println("Digite o salario do funcionário")
	fmt.Scan(&number)

	porcentagem := number * 0.15
	novo_salario := number + porcentagem
	fmt.Println("O funcionario que ganhava", number, "na promoção de 15% irá ganhar", novo_salario)

}

func Exerc14() {
	var number float64
	fmt.Println("Digite o valor da temperatura em Celsius")
	fmt.Scan(&number)
	fmt.Println(number)

	primeiro := number * 1.8
	resultado := primeiro + 32

	fmt.Println("Sua temperatura em Fahrenheit é de ", resultado)
}
