// [!] Hi, this demo is simple random variable checker
// [?] For the switch statement practice
// [*] 1 per 3 test some vars are equal!

package main

import (
    "fmt"
    "time"
    "math/rand"
)


func CheckIfNearNum(numToCheck int, conditionNum int,nameOfVar string) {
        switch numToCheck {
            case conditionNum:
                fmt.Printf("We got the success number -> %s = %d\n", nameOfVar, numToCheck)
            case conditionNum - 1, conditionNum + 1:
                fmt.Printf("We got the attempt to recive the number -> %d (warm)\n", numToCheck)
            default:
                fmt.Printf("We got no attempt to recive the number -> %d (cold)\n", numToCheck)
        }
}


func main() {
    for test := 0; test <= 10; test++ {
        var numToCheck, conditionNum int
        rand.Seed(time.Now().UnixNano())
        numToCheck = rand.Intn(10)
        conditionNum = rand.Intn(10)
        fmt.Println("-----------")
        CheckIfNearNum(numToCheck, conditionNum,"numToCheck")
        fmt.Print("[",numToCheck," ? ",conditionNum,"]\n")
    }
}
