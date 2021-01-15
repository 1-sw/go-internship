package main

import "fmt"
import "mask"

type Bit struct {
  Value bool
}

type Byte struct {
  Bits []Bit
}

type IP struct {
  NetworkAddr []Byte
  Hosts []Byte
  Port int
}

func AND(r int,l int) int {
  if(r == i && r == 1) {
    return 1
  } else {
    return 0
  }
}

func ConvertToBitFromBool(b bool) int {
    if(b) {
      return 1
    } else {
      return 0
    }
}

func ReadBytesFromBool(byt []Byte) {
  var bitCounter int
  for _,ByteValue := range byt {
    if(bitCounter >= 8) {
      fmt.Print(".")
      bitCounter = 0
    }
    for _,bit := range ByteValue.Bits {
      bitCounter++
      fmt.Print(ConvertToBitFromBool(bit.Value))
    }
  }
  fmt.Println("")

}

func main() {

  b0 := Bit{ Value: false,}
  b1 := Bit{ Value: true,}

  by11 := Byte{ Bits: []Bit { b0,b1,b1,b1, },}
  by12 := Byte{ Bits: []Bit { b1,b1,b1,b1, },}
  by21 := Byte{ Bits: []Bit { b0,b0,b0,b0, },}
  by22 := Byte{ Bits: []Bit { b0,b0,b0,b0, },}
  by31 := Byte{ Bits: []Bit { b0,b0,b0,b0, },}
  by32 := Byte{ Bits: []Bit { b0,b0,b0,b0, },}
  by41 := Byte{ Bits: []Bit { b0,b0,b0,b0, },}
  by42 := Byte{ Bits: []Bit { b0,b0,b0,b1, },}

  localhost := IP {
    NetworkAddr: []Byte{by11,by12},
    Hosts: []Byte{by21,by22,by31,by32,by41,by42},
    Port: 8080,
  }

  fmt.Println("-----------LOCALHOST-----------")
  fmt.Println("-----------SUBNET--------------")
  ReadBytesFromBool(localhost.NetworkAddr)
  fmt.Println("-----------HOSTS---------------")
  ReadBytesFromBool(localhost.Hosts)
  fmt.Println("-------------------------------")

}
