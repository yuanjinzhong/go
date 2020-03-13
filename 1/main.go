package main

import "fmt"

func main() {
	fmt.Println("hello world")
	var name string = "张三"
	var name1 = "李四"
	name2 := "王五"
	fmt.Println(name, name1, &name2)

	const age = 19
	fmt.Println(age)

	const a, b, c = age, "张三", false
	fmt.Println(a, b, c)

	var area int
	const len, withd = 10, 2
	area = len * withd
	fmt.Println(area)

	/*做枚举使用*/
	const (
		unknow = 0
		man    = 2
		femal  = 3
	)
	fmt.Println(unknow, femal, man)
}
