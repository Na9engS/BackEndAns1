package main

import (
	"fmt"
)

func drawPattern(x int) {
	// วาดดาวเพิ่มขึ้น
	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// วาดดาวลดลง
	for i := x - 1; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var x int
	fmt.Print("ใส่ค่าของ x: ")
	fmt.Scan(&x)
	drawPattern(x)
}
