package main

import "fmt"

type Phone interface {
   call()
}

type NokiaPhone struct {

}

func (nokia NokiaPhone) call()  {
  fmt.Println("i am nokia ")
}

type Iphone struct {
  
}

func (iphone Iphone)call()  {
 fmt.Println(" i am  iphone")
}


func main() {

  var  phone  Phone

  phone =new(NokiaPhone)
  phone.call()

  phone=new(Iphone)
  phone.call()
}
