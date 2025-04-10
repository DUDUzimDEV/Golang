package main 
  
  import "fmt"
 
  func main () {
 	/*ages := 16
    fmt.Println(ages <= 50)
    fmt.Println(ages >= 50)
    fmt.Println(ages == 50)
    fmt.Println(ages != 50)

    if ages < 15 {
        fmt.Println("Menor que 15 anos")
    } else if ages < 40 {
        fmt.Println("Menor que 40 anos")
    } else {
        fmt.Println("Não é menor que 40 anos")
    } */

    names := []string{"Isadora", "Yasmim", "Trunks", "Cebolinha", "Martin", "Yuri"}

    for index, value := range names {
        if index == 1 {
            fmt.Println("Continue após a posição", index, "e valor", value)
            continue
        }
        if index > 2 {
            fmt.Println("Sair após", index)
        }

        fmt.Println("Valor", value)
    }
  }