package main

import (
	"bytes"
	"fmt"
	"log"
)

// Ltime                         // the time in the local time zone: 01:23:23
// Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
// Llongfile                     // full file name and line number: /a/b/c/d.go:23
// Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
// LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
// LstdFlags

func examLogger() {
	var (
		buf     bytes.Buffer
		logger0 = log.New(&buf, "logger: ", log.Ltime)
		logger1 = log.New(&buf, "logger: ", log.Lmicroseconds)
		logger2 = log.New(&buf, "logger: ", log.Llongfile)
		logger3 = log.New(&buf, "logger: ", log.Lshortfile)
		logger4 = log.New(&buf, "logger: ", log.LUTC)
		logger5 = log.New(&buf, "logger: ", log.LstdFlags)
		logger6 = log.New(&buf, "logger: ", log.LstdFlags|log.Lshortfile)
	)

	logger0.Print("------------------")
	logger1.Print("------------------")
	logger2.Print("------------------")
	logger3.Print("------------------")
	logger4.Print("------------------")
	logger5.Print("------------------")
	logger6.Print("------------------")
	// logger0.Fatal("------------------Fatal------------------")
	fmt.Print(&buf)
}
