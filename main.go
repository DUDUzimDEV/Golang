package main 

import (
    "fmt"
    //"strings"
    "sort"
)

func main () {
    /*   greeting := "Hello my friends!"
    fmt.Println(strings.Contains(greeting, "friends"))
    fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
    fmt.Println(strings.ToUpper(greeting))
    fmt.Println(strings.Index(greeting, "my"))
    fmt.Println(strings.Split(greeting, "!"))    */
    
    ages := []int {50, 80, 10}
    sort.Ints(ages)
    fmt.Println(ages)
    index := sort.SearchInts(ages, 50)
    fmt.Println(index)
    names := []string{"Pedro", "Bruno", "Yan"}
    sort.Strings(names)
    fmt.Println(names)
    fmt.Println(sort.SearchStrings(names, "Pedro"))  
    
    
   /* x := 0
    for x <= 5 {
        fmt.Println(x)
        x++
    }

    for i := 0; i < 5; i++ {
        fmt.Println("for 2: ", i)
    }

    for i := 0; i < len(names); i++ {
        fmt.Println(names[1])
    }

    for index, value := range names {
        fmt.Println("O indice é: ", index, "e o valor ", value)
    } */

    for index, value := range ages {
        fmt.Println("O indice é:", index, "E o valor é: ", value)
    }

    superherois := []string{"DeadPool", "Homem Aranha", "Batman"}

    /*  for index, value := range superherois {
        fmt.Println("O número do herói é: ", index, "O nome do herói é: ", value)
    }   */

    for i:=0; i < len(superherois); i++ {
        fmt.Println("O número do herói é:", i, "e o herói é o ", superherois[1])
    }
}