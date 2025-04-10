package main 
  
  import "fmt"
 
  func main () {
 	var numeros = [5]int{}
 	fmt.Println("Digite 5 numeros")
    fmt.Scan(&numeros[0], &numeros[1], &numeros[2], &numeros[3], &numeros[4])
    fmt.Println("A soma dos numeros é: ", numeros[0] + numeros[1] + numeros[2] + numeros[3] + numeros[4] )
  }