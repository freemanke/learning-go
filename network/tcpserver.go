package main

import (
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Println("error listen:", err)
	}

	defer l.Close()
	var i int
	for {
		time.Sleep(time.Second * 1)
		if _, err := l.Accept(); err != nil {
			log.Println("accept error: ", err)
		}
		i++
		log.Printf("%d: accept a new connection\n", i)
	}
}
