package main

import (
	"fmt"
)

func FmtEntry() {
	examFormat()
}

type Person struct {
	name string
	age  int8
}

func examFormat() {
	strValue := "JKLMNOPQRST1234567"
	var byteValue byte = 0x0A
	personValue := Person{"Edwin", 28}
	personArray := []Person{Person{name: "&*^\"EdwinK", age: 22},
		Person{name: `@#$"EdwinW"`, age: 23}}
	_ = personArray
	_ = personValue
	fmt.Printf("character represtation c:%c\n", byteValue)
	fmt.Printf("value default format v:%v\n", byteValue)
	fmt.Printf("value default go format #v:%#v\n", byteValue)
	fmt.Printf("type T :%T\n", byteValue)
	fmt.Printf("%%十进制 d：%d\n", byteValue)
	fmt.Printf("二进制 b：%b\n", byteValue)
	fmt.Printf("八进制 o：%o\n", byteValue)
	fmt.Printf("十进制 d：%d\n", byteValue)
	fmt.Printf("十六进制 x：%x\n", byteValue)
	fmt.Printf("十六进制大写 X：%X\n", byteValue)
	fmt.Printf("%%十进制 #d：%#d\n", byteValue)
	fmt.Printf("二进制 #b：%#b\n", byteValue)
	fmt.Printf("八进制 #o：%#o\n", byteValue)
	fmt.Printf("十进制 #d：%#d\n", byteValue)
	fmt.Printf("十六进制 #x：%#x\n", byteValue)
	fmt.Printf("十六进制大写 #X：%#X\n", byteValue)
	fmt.Printf("单引号 q：%q\n", byteValue)
	fmt.Printf("Unicode U:%U\n", byteValue)
	fmt.Printf("------------STRING------------\n")
	fmt.Printf("String原始值 v：%v\n", strValue)
	fmt.Printf("String未解析原始值 s：%s\n", strValue)
	fmt.Printf("Double quoted q：%q\n", strValue)
	fmt.Printf("Base 16 Lower x：%x\n", strValue)
	fmt.Printf("Base 16 Upper X：%X\n", strValue)

	fmt.Printf("------------Slice&Pointer------------\n")
	strSlice := []byte(strValue)
	slicePointer := &byteValue
	fmt.Printf("Slice 0th 元素地址 p：%p\n", strSlice)
	fmt.Printf("Golang 原始值 v：%v\n", strSlice)
	fmt.Printf("Slice string未解析原始值 s：%s\n", strSlice)
	fmt.Printf("Double quoted q：%q\n", strSlice)
	fmt.Printf("Base 16 Lower x：%x\n", strSlice)
	fmt.Printf("Base 16 Upper X：%X\n\n", strSlice)
	fmt.Printf("Pointer 地址 p：%p\n", slicePointer)
	fmt.Printf("Pointer 地址 #p：%#p\n", slicePointer)
	fmt.Printf("Pointer 指向值0x0A v：%v\n", *slicePointer)
	fmt.Printf("Pointer 指向值0x0A b：%b\n", *slicePointer)
	fmt.Printf("------------Struct------------\n")
	ptrPerson := &personArray[0]
	fmt.Printf("Person对象默认格式 v：%v\n", personArray[0])
	fmt.Printf("Person对象默认格式 #v：%#v\n", personArray[0])

	fmt.Printf("Person对象指针默认格式 v：%v\n", ptrPerson)
	fmt.Printf("Person对象指针默认格式 #v：%#v\n", ptrPerson)

	fmt.Printf("------------Array------------\n")
	ptrPersonArr := &personArray
	fmt.Printf("Person数组默认格式 v：%v\n", personArray)
	fmt.Printf("Person数组默认格式 #v：%#v\n", personArray)
	fmt.Printf("Person数组指针默认格式 v：%v\n", ptrPersonArr)
	fmt.Printf("Person数组指针默认格式 #v：%#v\n", ptrPersonArr)

	fmt.Printf("------------Map&Pointer------------\n")
	psnMap := make(map[string]Person)
	psnMap["0"] = personArray[0]
	psnMap["1"] = personArray[1]
	ptrPsnMap := &psnMap
	fmt.Printf("PersonMap默认格式 v：%v\n", psnMap)
	fmt.Printf("PersonMap默认格式 #v：%#v\n", psnMap)
	fmt.Printf("PersonMapPointer默认格式 v：%v\n", ptrPsnMap)
	fmt.Printf("PersonMapPointer默认格式 #v：%#v\n", ptrPsnMap)
	type Car struct {
		engineSize int
		wheelCout  int
	}
	car1 := Car{engineSize: 1, wheelCout: 4}
	_ = car1
	car2 := Car{2, 4}
	carArr := []Car{{3, 3}, {4, 4}}
	fmt.Println(car1, car2, carArr)
	fmt.Printf("car1 %v, car2 %v, carArray %v\n", car1, car2, carArr)
	fmt.Printf("car1 %#v, car2 %#v, carArray %#v\n", car1, car2, carArr)
	// car2 :=

	// fmt.Printf("Double quoted：%q\n", strValue)
	// fmt.Printf("Base 16 Lower：%x\n", strValue)
	// fmt.Printf("Base 16 Upper：%X\n", strValue)
	// fmt.Printf("十进制：%d\n", byteValue)
	// fmt.Printf("十进制：%d\n", byteValue)
}
