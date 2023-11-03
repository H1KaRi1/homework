package main

import "fmt"

func Prime(n, a int) int {
	a = 0
	for i := 2; i < n; i++ {
		if n%i == 0 {
			a++
		}
	}
	return a
}
func main() {
	var n int
	fmt.Print("输入一个整数:")
	fmt.Scan(&n)
	var a int
	a = Prime(n, a)
	if a == 0 {
		fmt.Println("是素数")
	} else {
		fmt.Println("不是素数")
	}
}
