// Hi, I dont understand, wats going on

package main

import "fmt"

func main() {
    //so, we got the string:
    var anyText string = "Hello!"
    //and when we want to print symbol
    fmt.Println("Don't works: ",anyText[0])
    //It will print the byte of char: 70
    //But, if we use slices:
    fmt.Println(anyText[:1])
    //It works: H
}
