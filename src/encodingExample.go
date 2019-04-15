package main

import (
	"encoding/ascii85"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func examEncodingEntry() {
	// examAscii85()
	examAscii85Stream()
	// examDrivers()
}
func examAscii85() {
	encodingSrc := []byte("AAAAAAAAAAAAAAAAAa")

	encodingDst := make([]byte, ascii85.MaxEncodedLen(len(encodingSrc)), 100)
	encodedCount := ascii85.Encode(encodingDst, encodingSrc)
	fmt.Println("Total encoded length:", encodedCount)
	fmt.Println(encodingSrc)
	fmt.Println(encodingDst)
}
func examAscii85Stream() {
	os.Chdir("/Users/edwinking/Edwin/project001")
	myCmd, _ := exec.Command("pwd").Output()
	fmt.Println(string(myCmd))
	// println(os.Getenv("PWD"))
	myFile, err := os.Open("./ascii85.txt")
	if err != nil {
		fmt.Println(err)
		myFile, err = os.Create("./ascii85.txt")
	}
	defer myFile.Close()

	enCodeWriter := ascii85.NewEncoder(myFile)
	enCodeWriter.Write([]byte("AAAAAAAAAAAAAAAAAa"))
	// strArray := [100]byte{'B', 'B', 'B', 'B', 'B', 'B', 'B'}
	strArray := []byte("BBBBBBBBBB")
	strSlice := strArray[:]
	str := string(strSlice)
	_ = str
	fmt.Println(strSlice)
	str = strings.ToLower(str)
	strSlice = []byte(str)
	fmt.Println(strSlice)
	enCodeWriter.Close()

}
