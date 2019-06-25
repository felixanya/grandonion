package interface_demo

import "fmt"

type IReader interface {
	read()
}

type IWriter interface {
	write()
}

type RWriter struct {

}

func (rw *RWriter) read() {
	fmt.Println("do the read")
}

func (rw *RWriter) write() {
	fmt.Println("do the write")
}

func (rw *RWriter) say()  {
	fmt.Println("say something")
}

// 组合接口
type IReadWriter interface {
	IReader
	IWriter
}

type IReadWriterV2 interface {
	IReader
	IWriter
}

// test
func DoTest() {
	rw := &RWriter{}
	rw.say()

	var rw1 IReadWriter = rw
	rw1.read()
	rw1.write()
	//rw1.say()

	var rw2 IReadWriterV2 = rw
	rw2.read()
	rw2.write()
}
