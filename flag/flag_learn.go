package main

import (
	"flag"
	"fmt"
)

//flag包学习
var mode = flag.String("mode", "默认参数值", "玩玩命令行")

//玩玩命令行， 传启动参数go run flag_learn.go --mode=fast

//就像那些linux命令
/*$ go run flag_learn.go -help
Usage of C:\Users\yuanjz\AppData\Local\Temp\go-build256883454\b001\exe\flag_learn.exe:
-mode string
玩玩命令行 (default "默认参数值")
exit status 2
*/

func main() {
	flag.Parse()
	fmt.Println(*mode)
}
