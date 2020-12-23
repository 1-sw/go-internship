package main

import (
    "fmt"
    "math/rand"
    "time"
)


func main() {
    //self called anonymous func
    func() {
        fmt.Println(randInt())
    }()
    //So, we can change visibility with {}
    {
        hidden := 1
        fmt.Println(hidden)
    }

    //Anonymous func in var
    callMe := func() {
        fmt.Println("I am an anonymous func which works from the var")
    }
    //this will work:
    {
        callMe()
    }

    //Also we can do staff like this:
    wowExample := PrintName("Go")
    wowExample()
}

func randInt() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(10)
}


func PrintName(name string) func() {
    return func() {
        fmt.Printf("Hello, %s and nice to meet You!",name)
    }
}

