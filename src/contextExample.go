package main

import (
	"bytes"
	"context"
	"fmt"
)

var buf *bytes.Buffer

func examContext() {
	buf = bytes.NewBuffer(make([]byte, 0))
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.

	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1

		go func() {
			for {
				// time.Sleep(time.Millisecond)

			loop:
				select {
				case <-ctx.Done():
					//fmt.Println("main thread terminated************")
					buf.WriteString("main thread terminated************\n")
					return // returning not to leak the goroutine
				case dst <- n:
					if n == 2 {
						//fmt.Println("Channel closing!!!!!!!!!!!!!!!!!!")
						buf.WriteString("Channel closing!!!!!!!!!!!!!!!!!!\n")
						close(dst)
						//fmt.Println("Channel closed!!!!!!!!!!!!!!!!!!")
						buf.WriteString("Channel closed!!!!!!!!!!!!!!!!!!\n")
					}
					break loop
					n++
				}

			}
		}()

		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer deferredCancel(cancel) // cancel when we are finished consuming integers
	fmt.Println("************************************")

	for n := range gen(ctx) {
		buf.WriteString(fmt.Sprintln(n))
	}

}

func deferredCancel(param func()) {
	// fmt.Println("Running defer----------------")
	buf.WriteString("Running defer----------------\n")
	param()
	output := string(buf.Bytes())
	fmt.Println(output)
}
