package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//tty:pair<type:*os.file,value:/dev/tty>

	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	//r:pair<type:*os.file,value:E:\Go\src>
	var r io.Reader = tty

	var w io.Writer
	////w:pair<type:*os.file,value:E:\Go\src>
	w = r.(io.Writer)
	w.Write([]byte("hello world\n"))

	defer tty.Close()
}
