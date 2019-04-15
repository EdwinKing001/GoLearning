package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	/*
	   参数解析出错时错误处理方式
	   switch f.errorHandling {
	     case ContinueOnError:
	       return err
	     case ExitOnError:
	       os.Exit(2)
	     case PanicOnError:
	       panic(err)
	     }
	*/

	//flagSet = flag.NewFlagSet(os.Args[0],flag.PanicOnError)
	flagSet = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	//flagSet = flag.NewFlagSet("xcl",flag.ExitOnError)
	verFlag   = flagSet.String("ver", "", "version")
	xtimeFlag = flagSet.Duration("time", 10*time.Minute, "time Duration")

	addrFlag = StringArray{}
)

func init() {
	flagSet.Var(&addrFlag, "a", "b")
}

func examFlag() {
	fmt.Println("os.Args[0]:", os.Args[0])
	fmt.Println("os.Args[0]:", os.Args[1])
	fmt.Println("os.Args[0]:", os.Args[2])
	fmt.Println("os.Args[0]:%v", os.Args[:])

	flagSet.Parse(os.Args[1:]) //flagSet.Parse(os.Args[0:])
	for i := 0; i <= flagSet.NArg(); i++ {

		fmt.Println(flagSet.Arg(i))
	}

	fmt.Println("当前命令行参数类型个数:", flagSet.NFlag())

	fmt.Println("\n参数值:")
	fmt.Println("ver:", *verFlag)
	fmt.Println("xtimeFlag:", *xtimeFlag)
	fmt.Println("addrFlag:", addrFlag.String())

	for i, param := range flag.Args() {
		fmt.Printf("---#%d :%s\n", i, param)
	}
}

type StringArray []string

func (s *StringArray) String() string {
	return fmt.Sprint([]string(*s))
}

func (s *StringArray) Set(value string) error {
	*s = append(*s, value)
	return nil
}
