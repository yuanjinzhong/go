package slice

import (
  "fmt"
  "reflect"
  "testing"
  "unsafe"
)

type student struct {
  name string
  age  int
  sex  string
}

func TestArraySlice(t *testing.T) {

  student1 := student{"张三", 19, "女"}
  t.Log("结构体student实例", student1)
  t.Log("结构体student类型", reflect.TypeOf(student1))
  student2 := new(student)
  student2.age = 2
  student2.name = "里斯"
  fmt.Println("new 出来的结构体的值", student2)
  arr1 := make([]int, 2, 3)
  arr := [2]int{2, 3}
  slice := []int{3, 4}
  t.Log(arr[0], arr[1])
  t.Log(reflect.TypeOf(arr))
  t.Log(reflect.TypeOf(slice))
  t.Log(reflect.TypeOf(arr1))
  t.Log(unsafe.Sizeof(slice))
  t.Log(unsafe.Sizeof(arr))
  t.Log(append(slice, 2))

  // 直接声明切片 和make方式创建切片的区别：

  var myArr = [...]int{1, 3, 4, 5, 6}
  var mySlice = myArr[2:5]
  for _, v := range myArr {
    fmt.Println("数组的值为:", v)
  }

  for _, v := range mySlice {
    fmt.Println("切片的值为:", v)

  }
  t.Log("整体输出数组的值", myArr)
  t.Log("整体输出切片的值", mySlice)

  t.Log("切片的 容量为:", cap(slice))
  t.Log("切片的长度为: ", len(slice))

  t.Log("the type of myslice", reflect.TypeOf(mySlice))
  t.Log("the type of myArr", reflect.TypeOf(myArr))

  t.Log("--------------'make'a way to create slice ---------------------------------------------------------")

  mySlice2 := make([]int, 2, 3)
  mySlice2[0] = 1
  mySlice2[1] = 2
  t.Log("slice value created by the way of make: ", mySlice2)
  t.Log("slice capacity created by the way of make: ", cap(mySlice2))
  t.Log("slice len created by the way of make :", len(mySlice2))
  t.Log("slice type created by the way of make :", reflect.TypeOf(mySlice2))

}
