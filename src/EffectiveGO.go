package main

import "fmt"

func EffectiveGOEntry() {
	effectiveExample()
}

func effectiveExample() {
	v := 100
	fmt.Println(v, &v)
	for i := 100; i < 102; i++ {
		v := 101

		fmt.Println("v info:", v, &v)
		fmt.Println(i, &i)
		v = v + 1
	}
	fmt.Println(v, &v)

}
