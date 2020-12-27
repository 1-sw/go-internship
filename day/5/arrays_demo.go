package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//BASICS
	fmt.Println("1")
	test1()
	test2()
	test3()
	test4()
}

func test1() {

	fmt.Println("1.1")
	var arr1 = new([5]string)
	arr1[0] = "H"
	fmt.Println(arr1) //or:fmt.Println(&arr1)

	fmt.Println("1.2")
	var arr2 = new([1]float64)
	arr2[0] = 123.2222222
	fmt.Println(arr2) //or:fmt.Println(&arr2)


	fmt.Println("1.3")
	arr3 := []int{1,1,1,1}
	fmt.Println(arr3) //or:fmt.Println(&arr2)

}
func test2() {
	arr1 := []struct {
		name string
		surname string
	}{
		{"asdasd","zxczxc"},
		{"sddsas","qqwwqq"},
		{"asdasd","zxczxc"},
		{"sddsas","qqwwqq"},	
	}

	fmt.Println(arr1)

}
func test3() {
	arr1 := [5][5]byte {
		{1,2,3,4,5},
		{1,2,3,4,5},
		{1,2,3,4,5},
		{1,2,3,4,5},
		{1,2,3,4,5},
	}
	fmt.Println(len(arr1))
	for x := 0; x < len(arr1); x++ {
		for y := 0; y < len(arr1); y++ {
			fmt.Print(arr1[x][y])
		}
		fmt.Println("")
	}
}


func genXYZmatrix(x int, y int, z int){
	arr := [x][y][z]int{}
	for X := x; X < x; X++ {
		for Y := y; Y < y; Y++ {
			rand.Seed(time.Now().UnixNano())
			for Z := z; Z < z; Z++ {
				arr[X][Y][Z] := rand.Intn(123)
			}
		}
	}
	return arr
}

func test4(){
	result := genXYZmatrix(15,15,15)
	fmt.Println(result)
}