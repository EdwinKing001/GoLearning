package main

import (
	"fmt"
	"os"
	"strings"
)

func deffunc(file *os.File) {
	println("Close file")
	file.Close()
}

func examOSFile() {
	file, err := os.Open("edwin.txt")
	if err != nil {
		return
	}
	defer deffunc(file)
	defer func() string {
		println("Edwin")
		return ""
	}()
	//	file.Chdir()
	println(file.Name())
	readbuf := make([]byte, 100, 101)
	var byteRd int
	byteRd, err = file.ReadAt(readbuf, 6)
	if byteRd > 0 {
		println(string(readbuf))
		println("Bytes READ:", byteRd)
	}
	for err == nil {
		byteRd, err = file.Read(readbuf)
		if byteRd > 0 {
			print(string(readbuf))
		}
	}
	file.Close()
	file, err = os.Open(".//")
	fileinfoArray, errInfo := file.Readdir(5)
	if errInfo == nil {
		for _, fileinfo1 := range fileinfoArray {
			fmt.Println(fileinfo1)
		}
	}
	if err != nil {
		println(err.Error())
	}
}

func examOSPrintEvn() {
	fmt.Println("-------------Begin OS Example-------------")
	defer fmt.Println("-------------End   OS Example-------------")
	envMap := make(map[string]string)

	for _, osEnv := range os.Environ() {
		keyStr, valueStr := strings.Split(osEnv, "=")[0], strings.Split(osEnv, "=")[1]
		envMap[keyStr] = valueStr

	}
	fmt.Print("__________________________________________\n")
	fmt.Println(os.ExpandEnv("LogName:$LOGNAME"))
	fmt.Println(os.Getgid())
	fmt.Println(os.Getegid())
	fmt.Println(os.Geteuid())
	fmt.Println(os.Getwd())

	destStr := os.Getenv("GOHOME")
	fmt.Println(destStr)

	destVal, _ := os.Stat("./debug")
	interfaceValue := destVal.Sys()
	fmt.Println(interfaceValue)
	ret := fmt.Sprintf("%T", interfaceValue)
	ret += "Hello"
	fmt.Println(os.Lstat("./debuglink"))
	println(ret, ret)

}
