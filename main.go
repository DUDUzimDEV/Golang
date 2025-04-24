package main 
    
    import (
      "fmt"
    )

    var media float64
    var nota1 float64
    var nota2 float64

    func analisarNotas() (float64, string) {
      fmt.Println("Quais são as notas?")
      fmt.Scan(&nota1)
      fmt.Scan(&nota2)
      media := (nota1 + nota2) / 2
      if media >= 6 {
        return media, "Aprovado"
      } else if media <= 5 || media > 0 { 
        return media, "Reprovado"
      } else {
        return media, "Erro"
      }
    }
    
    func main () {
      media, resultado := analisarNotas()
      fmt.Println(media)
      fmt.Println(resultado)
    }