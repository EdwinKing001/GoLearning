package subpackage

import "fmt"

func Submain() {
	fmt.Println("------------------------------START------------------------------\n")
	// FmtEntry()
	// examFlag()
	// examEncodingEntry()

	// examDB_entry()
	// examCrypto()
	// baseAndPointer()
	// examgzip()
	// examContext()
	//examFlate()
	// examUnzip()
	//examLogger()
	// examOSFile()
	//examClosure1()
	//examClosure()
	// examByteReader()
	// examByte()
	//	examBuiltin()
	// examOSPrintEvn()
	//	examClosechannel()
	// examTim e()
	fmt.Println("------------------------------END------------------------------")
}

func typeTest() {
	intArray := [][]int{{1, 3, 5}, {2, 1, 6, 6, 3, 4}}
	fmt.Println(len(intArray[0]))
	fmt.Println(len(intArray[1]))
	fmt.Println(len(intArray))
	type tpitf interface{}
	var itf tpitf
	itf = int64(123)
	switch itf.(type) {
	case int:
		println("thi is int")
	case int64:
		println("thi is int64")
		println(itf.(int64))
	default:
		println("”UNKNOWN")
	}

}
