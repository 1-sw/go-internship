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


type ByteMethods interface {

    GetAllBits() []Bit
    GetBitByIndex(index int) Bit

    PushBit()
    PopBit()

}

type Byte struct {
    Bits []Bit
}

func (ByteV Byte) GetAllBits() []Bit {
    return ByteV.Bits
}

func (ByteV Byte) GetBitByIndex(index int) Bit {
    return Bit{
        _Value_ : bool(ByteV.Bits[index].GetValue()),
    }
}

func (ByteV Byte) PushBit() {
    fmt.Println("In process")
}

func (byteV Byte) PopBit() {
    fmt.Println("In process")
}

func CheckNewByte(ByteV ByteMethods) {
    fmt.Println("Maded new Byte")
}

func main() {

    var bit Bit = Bit {
        _Value_ : true,
    }

    var ByteV Byte = Byte {
        Bits : []Bit{bit},
    }

    CheckNewBit(bit)
    CheckNewByte(ByteV)
}
