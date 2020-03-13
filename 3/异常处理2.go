package main

import ("fmt"
"reflect")

type DivideError struct {
  dividee int
  divider int
}
type YuanJinz struct {
  name string
  age byte
}

func (de *DivideError) Error() string {
  strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
  return fmt.Sprintf(strFormat, de.dividee)
}

/*定义int 类型的除法运算函数*/
func Divide(varDividee int, varDivider int) (result int, errMsg string) {
  if varDivider == 0 {
    dData := DivideError{varDividee, varDividee}
    errMsg = dData.Error()
    return
  } else {
    return varDividee / varDivider, ""
  }
}

func main() {
  YuanJinz2:=new(YuanJinz)
  YuanJinz2.name="Azhon"
  YuanJinz2.age=255
  fmt.Println(YuanJinz2.name)
  fmt.Println(reflect.TypeOf(YuanJinz2))

  yjz:=YuanJinz{"yjz",28}
  fmt.Println(yjz.name)
  fmt.Println(reflect.TypeOf(yjz))

    //正常情况
    if result, errMsg := Divide(100, 10); errMsg == "" {
      fmt.Println("100/10结果:", result)
    }
    //异常情况
    if _, errorMsg := Divide(100, 0); errorMsg != "" {
      fmt.Println("errorMsg is", errorMsg)
    }
}