package uds

import (
	"io"
	"log"
	"net"
	"os"
)

const SocketAddr = "/tmp/echo.sock"

func echoServer(c net.Conn) {
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
	io.Copy(c, c)
	c.Close()
}

func ListenAndServe() {
	if err := os.RemoveAll(SocketAddr); err != nil {
		log.Fatal(err)
	}
	//rpc.Register()
	//rpc.Accept()
	server, err := net.Listen("unix", SocketAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := server.Close(); err != nil {
			log.Println("close server error, ", err.Error())
		}
	}()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go echoServer(conn)
	}
}
