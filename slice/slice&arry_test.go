package slice

import (
	"testing"
	"reflect"
	"unsafe"
	"fmt"
)

func TestArraySlice(t *testing.T)  {
	arr1:=make([]int,2,3)
	arr :=[2]int{2,3}
	slice :=[]int{3,4}
	t.Log(arr[0],arr[1])
	t.Log(reflect.TypeOf(arr))
	t.Log(reflect.TypeOf(slice))
	t.Log(reflect.TypeOf(arr1))
	t.Log(unsafe.Sizeof(slice))
	t.Log(unsafe.Sizeof(arr))
	t.Log(append(slice, 2))

	// 直接声明切片 和make方式创建切片的区别：

	var myArr=[...]int{1,3,4,5,6}
	var mySlice=myArr[2:5]
	for _,v:=range myArr{
		fmt.Println("数组的值为:",v)
	}

	for _,v:=range mySlice {
		fmt.Println("切片的值为:",v)

	}
	t.Log("整体输出数组的值",myArr)
	t.Log("整体输出切片的值",mySlice)

	t.Log("切片的 容量为:",cap(slice))
	t.Log("切片的长度为: ",len(slice))




}
