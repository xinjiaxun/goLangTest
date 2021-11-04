package main

import "fmt"
//全局变量
var name string = "pprof.cn"  //go 定义变量
var sex int = 1


func main ()  {

	age :=16  	//局部变量
	println(age+sex)
	println(name)
	println(sex)

	//数组
	a1 := [3] int {1, 2, 3}
	arr1 := [5] string {"1","2","3","4","5"}

	fmt.Println(a1)
	fmt.Println(arr1)

}
