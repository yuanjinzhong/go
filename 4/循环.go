package main

import (
  "fmt"
)

func main() {
  fmt.Println("loop learn")

  /* for true  {
     fmt.Println("this is a  endless loop")
   }*/

  for i := 5; i > 1; i-- {
    fmt.Println("now i is", i)
  }

  for j := 5; j > 1; j-- {
    for k := 5; k > 1; k-- {
      fmt.Println("the double loop result is ", k*j)
    }
  }
  name := "i am chinese"
  fmt.Println("name 的长度为", len(name))

  fmt.Println(add(4, 5))

  fmt.Println(swap("first", "second"))

}

func add(a, b int) int {
  var result = a + b;
  return result
}

func swap(x, y string) (string, string) {
  fmt.Println(" change argument")
  return y, x
}
