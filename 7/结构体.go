package main

import "fmt"

type Circle struct {
  radius float64
}

func (c Circle) getArea() (float64, string) {

  return 3.14 * c.radius * c.radius, "计算面积的结果"
}

func main() {

  var c1 Circle
  c1.radius = 10


  a,c:=c1.getArea()
  fmt.Println(a,c)



}
