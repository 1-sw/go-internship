//Matrix generator demo
//[+]every go programm consist of packages
//[+]every go programm starts with main
package main


import (
    "fmt"
    "time"
    "math/rand"
)


func genStringForm() string {
    forms := [] string { "|","-"," " }
    return forms[rand.Intn(len(forms))]
}

func genMatrix(x int, y int)  {
     rand.Seed(time.Now().UnixNano())
     form := genStringForm()
     for iter := 0; iter <= y; iter++ {
         for initer := 0; initer <= x; initer++ {
             fmt.Print(form,rand.Intn(9),form)
         }
         fmt.Printf("\n")
     }

}

func main() {
    for matrix := 0; matrix <= 10; matrix++ {
        genMatrix(matrix, matrix)
    }
}
