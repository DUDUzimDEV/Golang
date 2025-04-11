package main 
  
  import "fmt"

  var valorDEP float64
  var valorSAC float64
  var ValorSaldo float64
  var decisão float64
  var Decisão2 float64
 
  func saudar() {
    fmt.Println("Olá, tudo bem? O que você deseja fazer?")
    fmt.Println("Qual função você deseja realizar:\n1-Depositar\n2-Sacar\n3-Ver saldo")
    fmt.Scan(&decisão)
  }
  
  func sacar(){
    fmt.Println("Qual valor você deseja sacar?")
    fmt.Scan(&valorSAC) 
    if valorSAC < 0 || valorSAC > ValorSaldo {
        fmt.Println("Valor impossivel de ser sacado")
    } else {
        ValorSaldo = ValorSaldo - valorSAC
        fmt.Println("Seu novo saldo é ", ValorSaldo)
    } 
  }
 
  func depositar() {
    fmt.Println("Qual valor você deseja depositar?")
    fmt.Scan(&valorDEP)
    if valorDEP <= 0 {
        fmt.Println("Impossivel depositar numeros menores ou igual a 0")
    } else {
        ValorSaldo = ValorSaldo + valorDEP
        fmt.Println("Seu novo saldo é", ValorSaldo)
    }
  }
  
  func saldo() {
    fmt.Println("Seu saldo é ", ValorSaldo)
  }
  
  func Action() {
    if decisão == 1 {
        depositar()
    } else if decisão == 2 {
        sacar()
    } else if decisão == 3{
        saldo()
    } else {
        fmt.Println("Ação inexistente")
    }
  }


  func main () {
    fmt.Println("Qual é o seu saldo inicial?")
    fmt.Scan(&ValorSaldo)
    saudar()
    Action()
    for {
        fmt.Println("Você deseja realizar uma nova operação?\n1-Sim e 2-Não")
        fmt.Scan(&Decisão2)
        if Decisão2 == 2 {
            break
        } else if Decisão2 == 1{
            saudar()
            Action()
        } else {
            fmt.Println("Valor invalido")
        }
    }
  }