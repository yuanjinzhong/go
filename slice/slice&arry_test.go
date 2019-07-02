package slice

import (
	"testing"
	"reflect"
	"unsafe"
)

func TestArraySlice(t *testing.T)  {
	arr :=[2]int{2,3}
	slice :=[]int{3,4}
	t.Log(arr[0],arr[1])
	t.Log(reflect.TypeOf(arr))
	t.Log(reflect.TypeOf(slice))
	t.Log(unsafe.Sizeof(slice))
	t.Log(unsafe.Sizeof(arr))
	t.Log(append(slice, 2))

}
