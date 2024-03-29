package main

import "fmt"

// 定义一个结构体 person
type person struct {
	name string
	age  int
}

func main() {
	var p person   // 声明一个 person 类型变量 p
	p.name = "max" // 赋值
	p.age = 12
	fmt.Println(p) // 输出：{max 12}

	p1 := person{name: "mike", age: 10} // 直接初始化一个 person
	fmt.Println(p1.name)                // 输出：mike

	p2 := new(person) // new函数分配一个指针，指向 person 类型数据
	p2.name = `张三`
	p2.age = 15
	fmt.Println(*p2) // 输出：{张三 15}
}
