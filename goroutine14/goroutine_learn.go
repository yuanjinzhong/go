package main

import (
  "fmt"
  "runtime"
  "time"
)

var cpuNum=runtime.NumCPU()

func say(s string) {
  fmt.Println(s)
  time.Sleep(2 * time.Second)
}

func main() {
  fmt.Println("cpu核心数量",cpuNum)
  runtime.GOMAXPROCS(cpuNum)
  //defer 语句总是最后执行
  defer fmt.Println("dsdsd")
  go say("world")
  say("hello")
//若主函数推出，则goroutine 会立刻退出
}
