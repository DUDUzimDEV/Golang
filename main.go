package main 

import "fmt"

func main () {
 var num1 float32
 var num2 float32
 fmt.Println("Digite dois numeros:")	
 fmt.Scan(&num1, &num2)
 fmt.Println("A soma é: ", num1 + num2)
 fmt.Println("A subtração é: ", num1 - num2)
 fmt.Println("O multiplicação é: ", num1 * num2)
 fmt.Println("A divisão é: ", num1 / num2)
}