package main

import "fmt"

type Person struct{
	name , password string;
}
func say传值(p Person){
	p.name="李四"
}

func say传引用(p *Person){
	p.name="李四"
}

func main (){
	person:=Person{"张三","12345"}

	say传值(person)
	fmt.Print(person)

	personPoint:=new(Person)
	personPoint.name="张三"
	personPoint.password="asdf2580"
	say传引用(personPoint)
	fmt.Print(personPoint)

	person2:=Person{"张三","12345"}
	say传引用(&person2)
	fmt.Print(person2)




}
