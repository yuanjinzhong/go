package main

import "fmt"

//通道学习
//If zero, or the size is omitted, the channel is	unbuffered.
//make(chan int,8) 带8位字节的通道
//var chanel = make(chan int)  不带缓冲的

func sum(arr []int, c chan int) {
  sum := 0
  for _, v := range arr {
    sum += v
  }
  c <- sum
}

func main() {
  arr := [...]int{1, 2, 3, 4, 5}
  channel := make(chan int)
  go sum(arr[:3], channel)
  go sum(arr[4:], channel)

  x, y := <-channel, <-channel
  fmt.Println("2通道里面的结果为", x+y)
}
