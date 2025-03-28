package main 

import "fmt"

func main () {
 var usuario string
 var senha int
 fmt.Println("Digite seu nome de usuario: ")
 fmt.Scan(&usuario)
 if usuario == "admin" {
   fmt.Println("Digite a senha")
   fmt.Scan(&senha)
   if senha == 1234 {
	fmt.Println("Você fez login com sucesso")
   } else {
	fmt.Println("Senha errada")
   }
 } else {
	fmt.Println("Nome de usuario errado")
 }
}