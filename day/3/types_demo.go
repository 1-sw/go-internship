/*
                  numbers
                 /        \
          integers         floats
          /      \         /   | \
         uint*   int*   float32|complex64
 byte<->uint8   int8    float64|complex128
        uint16  int16          |
        uint32  int32<->rune  NaN
        uint64  int64      (not a number)

  *- int,uint,uintptr - machine influenced types,
   - because them size has dependences from architecture
   - of your machine
*/


package main

import "fmt"

type SomeOtherString string

func integers() {
    fmt.Println("Testing: INTEGERS")
    var (
        someUint8 uint8 = 1
        someByte  byte  = 2
        someInt32 int32 = 123123123
        someRune  rune  = 123123123
        someUintptr uintptr = 234
    )
    test1 := someUint8 + someByte  //But we cant mix any other not identical types, for example
    test2 := someInt32 + someRune  // someInt32 + someUint8 will throw an matched type error
    fmt.Println(test1,test2,someUintptr)
}

func floats() {
    fmt.Println("Testing: FLOATS")
    var (
        someFloat32 float32 = 2123.123123
        someNum = 2123.123123
        someFloat64 float64 = 2123.123123
    )
    func () {
        if(someNum == someFloat64) {
            fmt.Println(someFloat32)
        }
    }()
}

func main() {
    integers()
    floats()
}
