package main

import (
    "fmt"
    "time"
)


func printByCount(data string, count byte) {
    for iter := byte(0); iter <= count; iter++ {
        fmt.Printf(data)
    }
    time.Sleep(time.Second)
}

func endl() {
    fmt.Printf("\n")
}

func main() {
    for routineCycle := 0; routineCycle <= 3; routineCycle++ {
        fmt.Printf("[Routine cycle: %s]\n", routineCycle)
        go printByCount("r3 -> ",1)
        endl()
        go printByCount("r2 -> ",1)
        endl()
        printByCount("r1 ->",1)
        endl()
    }
}
