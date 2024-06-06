package main

func main() {
	b := [...]int{109, 110, 111}
	p := &b
	p++
}

//上面的代码执行将报错，因为指标不支持运算
