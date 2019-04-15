package main

import "fmt"

type COLOR int
type CarType int

const (
	TRUCK CarType = iota
	CAR
	MOTOR
	CRANE
)
const (
	WHITE COLOR = iota
	BLACK
	BLUE
	RED
)

func EffectiveGOEntry() {
	newmakeExample()
}

type car struct {
	color COLOR
	model string
	ctype CarType
}

// func (car) String() string {
// 	return "this is a car"
// }

func (*car) String() string {
	return "this is a car"
}

func newmakeExample() {
	//初始化结构体是，各个字段之间使用comma分割
	car1 := car{color: RED, model: "falali", ctype: CAR} // 使用冒号给特定的field赋值
	car1 = car{color: WHITE}                             //结构体的部分赋值，未指定部分初始化为zero-value

	//不指明字段名称是，必须给全部的field初始值
	car1 = car{BLUE, "HYNDA", TRUCK}

	//注意下面的双层brace，数组字面初始化时，使用brace而不是中括号
	carArray := []car{{BLUE, "AUDI", CRANE}, {BLUE, "VOLKS-WAGON", MOTOR}}
	carPtr := new(car)

	// 使用backslash转移字符
	fmt.Println("\n", *carPtr)
	carPtr = &car{color: RED, model: "falali", ctype: CAR}
	fmt.Println(*carPtr)

	fmt.Println(car1)
	fmt.Println(carArray)
	_ = carPtr

	// 数组，slice，map通过给定的index和key来进行初始化，类似初始化一个struct
	a := [...]string{1: "no error", 2: "Eio", 5: "invalid argument"}
	s := []string{6: "no error", 9: "Eio", 55: "invalid argument"}
	m := map[int]string{4: "no error", 8: "Eio", 889: "invalid argument"}

	fmt.Println(len(a), "\n", len(s), "\n", len(m), "\n")
	fmt.Println(cap(a), "\n", cap(s), "\n")

}

func effectiveExample001() {
	v := 100
	fmt.Println("v Declard every thime the for loop runs, the address are always different.")
	fmt.Println(v, &v)
	for i := 100; i < 102; i++ {
		v := 101

		fmt.Println("v info:", v, &v)
		fmt.Println(i, &i)
		v = v + 1

	}
	fmt.Println(v, &v)

}
