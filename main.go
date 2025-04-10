package main 
  
  import "fmt"
 
  func main () {
 	var ages int
    fmt.Println("Digite sua idade")
    fmt.Scan(&ages)
    if ages < 18 {
        fmt.Println("Você é menor de idade")
    } else if ages > 18 && ages <= 60 {
        fmt.Println("Você é adulto")
    } else if ages >60 {
        fmt.Println("Você é idoso")
    } 
  }