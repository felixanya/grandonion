package uds

import (
	"log"
	"net"
)

func (u *UDS) ClientSendContext(context string) string {
	sockAddr, err := net.ResolveUnixAddr("unix", u.sockFile)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUnix("unix", nil, sockAddr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte(context))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, u.bufSize)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	return string(buf[0:n])
}
