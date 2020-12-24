package main

import (
	"fmt"
	//"math"
)


const (

	BIT1   byte = 0b1
	MAX1_2 byte = 0b11       // -> (3|256)
	MAX1_3 byte = 0b111      // -> (7|256)
	MAX1_4 byte = 0b111      // -> (15|256)
	MAX1_8 byte = 0b11111111 // -> (255|256)

	BIT0   byte = 0b0
	MAX0_2 byte = 0b00       //  -> (0|256)
	MAX0_3 byte = 0b000      //  -> (0|256)
	MAX0_8 byte = 0b00000000 //  -> (0|256)
)

func getByteNum(size string) byte {
	var varToReturn byte
	switch size {
	case "MIN": 
		varToReturn = BIT0
	}
	return varToReturn
}
func main() {
	fmt.Println(BIT0, MAX0_2, MAX0_3, MAX0_8) 
}