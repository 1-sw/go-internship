package main

import "fmt"

type BaseValue interface {
    SetTrueValue() bool
    SetFalseValue() bool
    GetValue() bool
}

type Bit struct {
    _Value_ bool
}

func (bit Bit) SetTrueValue() bool {
    bit._Value_ = true
    return bit._Value_
}

func (bit Bit) SetFalseValue() bool {
    bit._Value_ = false
    return bit._Value_
}

func (bit Bit) GetValue() bool {
    return bit._Value_
}

func CheckNewBit(bit BaseValue) {
    fmt.Println("Bit maded")
    fmt.Println("Some value:")
    fmt.Println(bit.GetValue())
}

func main() {
    var bit Bit = Bit {
        _Value_ : true,
    }
    CheckNewBit(bit)
}
