package main

import (
	"fmt"
	"time"
)

var outer = 0
var pointer *int = nil

func examClosure() {

	printInt := func(x int) string {
		// println("Simple closure:", x, " outer:", outer)
		outer += x

		inner := -1 * x

		return fmt.Sprintf("return: %d, outer:%d , inner:%d", x, &outer, &inner)
	}

	for index := 0; index < 10; index++ {
		println(printInt(100 + index))
		println(printInt(-10 - index))
	}
	println(*pointer)

}
func examClosure1() {
	var flist []func()
	for i := 0; i < 900000; i++ {

		var i1 int
		i1 = i1 //给i变量重新赋值，
		if pointer == nil {
			pointer = &i1
		} else {
			//println(*pointer)
		}
		i1 = i

		// fmt.Println(&i1)
		// i1 = ii
		//	fmt.Println(&i1, " ", &i)
		flist = append(flist, func() {
			// fmt.rintf("addr of i: %X value:", &i1)
			i1 = i1 * 100
			// fmt.Println(i1)

		})
	}
	for _, f := range flist {
		f()
	}
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() && false {
	// 	userInputText := scanner.Text()
	// 	if userInputText == " " {
	// 		return
	// 	} else {
	// 		println("userInputText")
	// 	}
	// }
	timechan := time.After(6 * time.Second)
	timeNow := time.Now()
	fmt.Println(timeNow)
	timeNow = <-timechan
	fmt.Println(timeNow)

}
