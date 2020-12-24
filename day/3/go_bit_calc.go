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

func getByteValueByCount(countOfSymbols uint8, type01 uint8) byte {
	switch type01 {
		case 0: 
			return BIT0
		case 1:
			switch countOfSymbols {
				case 2 :
					return MAX1_2
				case 3 :
					return MAX1_3				
				case 4 :
					return MAX1_4				
				case 8 :
					return MAX1_8							
			}
		default:
			return MAX0_8
	}
	return BIT0
}
func main() {
	fmt.Println(getByteValueByCount(2,1)) 
}