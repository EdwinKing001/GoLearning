package main

import "fmt"
import (
	"bytes"
	"time"
)

func examByteReader() {
	fmt.Println("------------------START------------------")

	var byteArr01 [100]byte
	byteSlice := byteArr01[:] //make([]byte, 0, 100)
	index := byte(0)
	for index < 100 {
		byteSlice[index] = 'A' + index%62
		index++
	}
	myReader := bytes.NewReader(byteSlice)

	for {
		ret, size, errRead := myReader.ReadRune()
		if errRead == nil {
			println(string(ret), size)
		} else {
			println(errRead.Error())
			return
		}

	}
}
func examByte() {
	fmt.Println("------------------START------------------")

	var byteArr01 [100]byte
	var byteSlice = byteArr01[:] //make([]byte, 0, 100)
	{
		stringFromArray := string(byteArr01[:])
		println(stringFromArray)
	}
	index := byte(0)
	for index < 100 {
		byteSlice[index] = 'A' + index%62
		index++

	}
	// ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~�������������������������������������

	string01 := string(byteSlice) //slice to string01
	//	byteArr02 := []byte(string01) // string to slice
	// println(string01)
	byteBuffer01 := bytes.NewBuffer(byteSlice)
	println(byteBuffer01.String())
	_, err := byteBuffer01.Write(byteSlice)
	if err != nil {
		println(err)
	}
	println(byteBuffer01.String())
	byteBuffer01 = bytes.NewBufferString(string01)
	println(byteBuffer01)
	fmt.Printf("Type is:%T\n", byteArr01)
	fmt.Printf("Type is:%T\n", byteSlice)
	fmt.Printf("Type is:%T\n", byteBuffer01)
	fmt.Printf("Type is:%T\n", string01)
	type i interface{}
	var i1 i = string01
	switch i1.(type) {
	case int:
		return
	case byte:
		return
	default:
		break
	}

	fmt.Println(time.Now(), "------------------JOB DONE------------------")

}
