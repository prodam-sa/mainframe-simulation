package main

import (
	"net"
	"log"
	"strconv"
)

// MainframeSocket struct
type MainframeSocket struct {
	Host string
	Port int
	Tipo string
	Count int
	Server net.Listener
	Lock chan int
}

func main() {
	var err error
	var lock chan int
	var mainframeSockets []*MainframeSocket

	for i := 0; i < 4; i++ {
		socket := MainframeSocket{
			Host: "localhost",
			Port: 3000 + i,
			Tipo: "tcp",
			Count: 0,
		}
		mainframeSockets = append(mainframeSockets, &socket)
	}

	for _, m := range mainframeSockets {
		socket := m
		go startServer(socket, err)
	}

	<-lock
}

func startServer(socket *MainframeSocket, err error) {
	socket.Server, err = net.Listen(socket.Tipo, socket.Host + ":" + strconv.Itoa(socket.Port))
	if err != nil {
		log.Println(err)
	}

	conn, err := socket.Server.Accept()
	if err != nil {
		log.Println(err)
	}

	handleRequest(conn, socket)
}

func handleRequest(conn net.Conn, socket *MainframeSocket) {
	var c []byte
  for {
		conn.Write([]byte("Mensagem recebida pelo servidor\n"))
		socket.Lock <- conn.Read(c)
	}
}
