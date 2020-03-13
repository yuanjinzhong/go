package main

import (
  "errors"
  "fmt"
)

/*获取平方根*/
func getSquareRoot(x int) (int, error) {
  if x < 0 {
    return 0, errors.New("数据不能为负数")
  }
  fmt.Println("这是正常实现")
  return 8, nil
}

func main() {
  result, err := getSquareRoot(-2)
  if err != nil {
    fmt.Println(err)

  } else {
    fmt.Println(result)
  }
}
