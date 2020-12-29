/*

Hi, this file was uploaded from previous commits,
I prefer to clear all files, because I need to much
time for logging them.

Next question:
    We have any string `text := string("hi!")`
    but when we deside to print one symbol:
    `fmt.Println(text[0])` it will print thy byte: 104
    Why?

Answer:
    string in go is the sequens of bytes, here the implementation:
    type _string struct {
       elements *byte
       len      int
    }
    So, we just call one of those bytes
    [0] h : 104
    [1] i : 105
    [2] ! : 33

Sources:
    https://blog.golang.org/slices-intro
    https://go101.org/article/string.html
    https://yourbasic.org/golang/convert-string-to-rune-slice/
*/



package main

import (
    "fmt"
)


func main() {

    text := string("hi!")
    textb:= byte(text[0])

    fmt.Printf("Var:text type:%T value:%s\n",text,text)
    fmt.Printf("Var:textb type:%T value:%s (converted -> %s)\n",textb,textb,string(textb))

    //for i:=0; i<len(text); i++ {
    //    fmt.Println(text[i])
    //}

}
