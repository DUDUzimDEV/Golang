package main 

import "fmt"

func main () {
	var numero int
	fmt.Print("Digite um número ")
	fmt.Scan(&numero)
	fmt.Println("O numero digitado", numero)
}