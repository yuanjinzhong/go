module go

go 1.14

require (
	github.com/astaxie/beego v1.12.1
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
)

//todo  引入本地包有问题
replace github.com/xxx => ../go_respority/go_utils/stringutil

