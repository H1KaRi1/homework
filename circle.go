package main

import "fmt"
import "math"

func Square(r float64) float64 {
	area := math.Pi * r * r
	return area
}
func main() {
	var r float64
	fmt.Print("请输入圆的半径:")
	fmt.Scan(&r)
	area := Square(r)
	fmt.Println("圆的面积是:", area)
}
