package main

import (
	"bufio"
	"fmt"
	"fyne-app/file/hello"
	"io/ioutil"
	"os"
)

func main() {
	wt := func(f *os.File, s string) {
		_, er := f.WriteString(s + "\n")
		if er != nil {
			panic(er)
		}
	}

	fn := "data.txt"

	fs, er := ioutil.ReadDir(".")
	if er != nil {
		panic(er)
	}

	for _, f := range fs {
		fmt.Println(f.Name(), "(", f.Size(), ")")
	}

	f, er := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if er != nil {
		panic(er)
	}

	fmt.Println("*** start ***")

	wt(f, "*** start ***")

	for {
		s := hello.Input("Type message...")

		if s == "" {
			break
		}
		wt(f, s)
	}

	wt(f, "*** end ***\n\n")

	fmt.Println("*** end ***")

	defer f.Close()

	// 読み込み
	rt := func(f *os.File) {
		r := bufio.NewReaderSize(f, 4096)

		for i := 1; true; i++ {
			s, _, er := r.ReadLine()
			if er != nil {
				break
			}
			fmt.Println(string(s))
		}
	}

	f2, er2 := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)

	if er2 != nil {
		panic(er2)
	}

	defer f2.Close()

	fmt.Println("<< start >>")
	rt(f2)
	fmt.Println("<< end >>")
}
