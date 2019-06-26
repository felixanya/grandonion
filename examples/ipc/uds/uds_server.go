package uds

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"time"
)

type UDS struct {
	sockFile 		string
	bufSize 		int
	handler 		func(string) string
}

func NewUDS(sockFile string, size ...int) *UDS {
	defaultSize := 10480
	if size != nil {
		defaultSize = size[0]
	}
	
	sock := &UDS{
		sockFile: sockFile,
		bufSize: defaultSize,
	}
	return sock
}

func (u *UDS) ListenAndServe() {
	if err := os.RemoveAll(u.sockFile); err != nil {
		log.Fatal(err)
	}
	sockAddr, err := net.ResolveUnixAddr("unix", u.sockFile)
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.ListenUnix("unix", sockAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := listen.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for {
		conn, err := listen.AcceptUnix()
		if err != nil {
			log.Println(err)
		}
		go u.HandleServerConn(conn)
	}
}

func (u *UDS) HandleServerConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, u.bufSize)

	// 循环接收socket数据
	// 但，Read()并不是阻塞的
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
		}

		recv := &RecvData{}
		if err := json.Unmarshal(buf[0:n], recv); err != nil {
			log.Printf("unmarshal msg error, %s\n", err.Error())
		}
		//log.Printf(">>>%+v \n", recv)

		result := u.HandleServerContext(string(buf[0:n]))
		_, err = conn.Write([]byte(result))
		if err != nil {
			log.Fatal("write error,", err)
		}
	}

}

func (u *UDS) SetContextHandler(f func(string) string) {
	u.handler = f
}

func (u *UDS) HandleServerContext(context string) string {
	if u.handler != nil {
		return u.handler(context)
	}
	now := time.Now().String()
	return now
}

func (u *UDS) StartServer() {
	u.ListenAndServe()
}

func main() {
	server := NewUDS("/tmp/echo.sock")
	server.StartServer()
}

type RecvData struct {
	Timestamp 		int 		`json:"timestamp"`
	Delay 			int 		`json:"delay"`
	Key 			int 		`json:"key"`
}
