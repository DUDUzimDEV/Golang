package main 
    
    import (
      "fmt"
    )

    func main () {
      notasAlunos := map[string]float64{
        "Bruno" : 9.7,
        "Otavio" : 10,
        "Fabiano" : 8.7,
        "Isabela" : 9.5,
      }
      for nome,nota := range notasAlunos {
        fmt.Printf("%s tirou a nota %.1f \n", nome, nota)
      }
    }