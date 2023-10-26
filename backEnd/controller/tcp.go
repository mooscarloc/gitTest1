package controller

import (
	"bufio"
	"data/models"
	"fmt"
	"io"
	"net"
)

func init() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("System recovered", r)
			}
		}()
		
		var tcpAddr *net.TCPAddr
		tcpAddr, _ = net.ResolveTCPAddr("tcp", ":8000")
		tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
		defer tcpListener.Close()
		fmt.Println("Server ready to read ...")
		for {
			tcpConn, err := tcpListener.AcceptTCP()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("A client connected :" + tcpConn.RemoteAddr().String())
			go tcpConnection(tcpConn)
		}
	}()

}

func tcpConnection(conn *net.TCPConn){
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println(" Disconnected : " + ipStr)
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	for {
		message,err := reader.ReadString('\n')
		if err!=nil || err == io.EOF {
			break
		}
		
		models.InsertRecord(message)
	}
}
