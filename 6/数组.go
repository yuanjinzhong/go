package main

import "fmt"

var age = [4]int{1, 2, 3, 4}

func main() {

	for i := len(age); i > 1; i-- {
		fmt.Println("当前数组下标", i)
	}
	fmt.Println("数组的地址是", *&age)

	var a int = 10

	fmt.Printf("变量的地址: %x\n", &a)

}
