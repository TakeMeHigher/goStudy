package 可见性
// 大写能被外部的包引用  小写 说明是私有的 只能在本包里面被调用 (本包里面的所有.go都可调用)

var a=100
var A = 200

func Add(a, b int) int {
	return a+b
}

func sub(a, b int) int {
	return a - b
}