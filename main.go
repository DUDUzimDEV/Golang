package main 
  
  import "fmt"
 
  func main () {
 	var saldo float64
    var decisão float64
    var sacado float64
    var depositado float64
    fmt.Println("Digite seu saldo")
    fmt.Scan(&saldo)
    fmt.Println("Digite 1 para Sacar e 2 para Depositar")
    fmt.Scan(&decisão)
    if decisão == 1 {
        fmt.Println("Qual valor você decide sacar?")
        fmt.Scan(&sacado)
        if sacado > saldo || sacado < 0 {
            fmt.Println("Valor indisponivel  ")
        } else {
            fmt.Println("Seu novo saldo é ", saldo - sacado)
        }
    } else {
        if decisão == 2 {
        fmt.Println("Qual valor você deseja depositar?")
        fmt.Scan(&depositado)
        if depositado < 0 {
            fmt.Println("Impossivel adicionar valores menores que 0")
        } else {
            fmt.Println("Seu novo saldo é", saldo + depositado)
        }
    } else {
        fmt.Println("Error")
    }
  }  }