package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func examIO() {
	r1 := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")
	b, err := ioutil.ReadAll(r1)
	if err != nil {
		log.Fatal(err)
	}
	//Go is a general-purpose language designed with systems programming in mind.
	fmt.Printf("%s\n len is:%d\n，cap is:%d\n", b, len(b), cap(b))

	r2, err1 := ioutil.ReadDir("/Users/edwinking/Edwin/project001")
	if err1 != nil {
		log.Fatal(err1)
	}
	for _, v := range r2 {
		// info := v
		println(v.Name())

		// if v == "debuglink" {
		//   os.Rename(v.Name(), newpath string)
		// }

	}

}
