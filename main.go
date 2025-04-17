package main 
    
    import (
      "fmt"
    )

    func infoPessoais(idade int, nome string) (int, string){
      if idade <= 0  || idade > 110 {
        fmt.Println("Idade invalida")
      } else if idade < 18 || idade > 0 {
        fmt.Println("Você é menor de idada") 
      } else {
        fmt.Println("Você é maior de idade")
      }
      return idade, nome
    }
    
    func main () {
      var idade int
      var nome string
      fmt.Println("Qual a sua idade?")
      fmt.Scan(&idade)
      fmt.Println("Qual o seu nome?")
      fmt.Scan(&nome)
      
      fmt.Println("Nome: ", nome)
      fmt.Println("Idade: ", idade)
      infoPessoais(idade, nome)
    }