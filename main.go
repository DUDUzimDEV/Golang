package main 
    
    import (
      "fmt"
    )

    func main () {
      inventario := map[string]int {
        "COXINHA": 10,
        "PÃO DE QUEIJO": 15,
        "REFRIGERANTE": 20,
      }

      for K, V := range inventario {
        fmt.Println("Item:", K, "Quantidade:", V)
      }
      
      fmt.Println("------------------------------------------------------")

      inventario["COXINHA"] -= 2

      inventario["PÃO DE QUEIJO"] -= 1

      for K, V := range inventario {
        fmt.Println("Item:", K, "Quantidade:", V)
      }
    }