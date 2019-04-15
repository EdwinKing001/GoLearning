package main

import (
	"fmt"
	"io"
)

type holder interface{}

func tryInterfaceType(itfc interface{}) {
	switch itfc.(type) {
	case chan int:
		println("int channel")
	case chan string:
		println("string channel")
	default:
		println("unknown type")

	}
}
func pushIntChannel(intChan chan int) {

}
func pushStringChannel(intChan chan string) {

}
func examBuiltin() {
	var intArr0 [32]int16
	var intArr1 [64]int32
	intSlice0 := intArr0
	intSlice1 := intArr1

	println("array cap example:", cap(intArr0), cap(intArr1))
	println("array len example:", len(intArr0), len(intArr1))

	println("--------------------------------------------------")
	println("slice cap example:", cap(intSlice0), cap(intSlice1))
	println("slice len example:", len(intSlice0), len(intSlice1))
	var i holder
	var makeslice = make([]int, 5, 10)
	println(makeslice)
	makemap := make(map[int]string)
	println(makemap)
	makechannel := make(chan int, 10)
	makechannel1 := make(chan string)
	i = makechannel
	tryInterfaceType(i)
	fmt.Printf("Type is %T\n", i)

	i = makechannel1
	tryInterfaceType(i)
	fmt.Printf("Type is %T\n", i)
	interPointer := new(int)
	*interPointer = 0xFFFF
	println(interPointer, *interPointer)
	// tryRecover()
}
func tryRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错：", err)
		}
	}()
	tryPainc(nil)
	tryPainc(io.ErrNoProgress)
	fmt.Printf("这里应该执行不到！")
}
func tryPainc(param error) {
	var x = 30
	var y = 0
	if param == nil {
		panic("我就是一个大错误！")
	} else if param != io.ErrUnexpectedEOF {
		panic(param)
	}
	var c = x / y
	fmt.Println(c)
}
