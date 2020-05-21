package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	conn, err := net.Dial("tcp" , ":2345")
	if err != nil {
		log.Println(err)
	}

	go func(){
		data := make([]byte ,4096)

		for {
			n,err := conn.Read(data)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("Sever send : " + string(data[:n]))
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()

	for {
		var s string
		fmt.Scanln(&s)
		conn.Write([]byte(s))
		time.Sleep(time.Duration(1) * time.Second)
	}
}