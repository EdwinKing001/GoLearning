package main

import "fmt"
import "time"

func examTime() {
	fmt.Println(time.Now(), "START**********************************")
	var tm time.Time

	tm1 := time.Now()
	for index := 0; index < 5; index++ {
		tm = <-time.After(50 * time.Millisecond)
		fmt.Println(tm, "timed out")
	}
	fmt.Println(tm1, "JOB DONE")

}
