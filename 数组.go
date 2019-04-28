package main

import "fmt"

var age [4] int

func main() {

  for i := len(age); i > 1; i-- {
    fmt.Println("当前数组下标", i)
  }
fmt.Println("数组的地址是",&age)

}
