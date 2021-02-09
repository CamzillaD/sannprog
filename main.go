package main

import (
	"fmt"
	"net"
	"time"
)

const (
	connHost = "localhost"
	connPort = "30000"
	connType = "udp"
)

func UDP_listen(port int) {

	pc, err := net.ListenPacket("udp4", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	buf := make([]byte, 1024)
	for {
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s sent this: %s\n", addr, buf[:n])
		time.Sleep(time.Second)
	}
}

func UDP_send() {

	addr, err := net.ResolveUDPAddr("udp4", "10.100.23.147:20016")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		panic(err)
	}

	for {
		conn.Write([]byte("hei"))
		time.Sleep(time.Second)
	}

}

func TCP_connect() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", "10.100.23.147:34933")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}

	go TCP_reciever(conn)

	for {
		conn.Write(append([]byte("hallo jeg er her elns xD lol\n\r"), 0))
		time.Sleep(time.Second)
	}
}

func TCP_reciever(conn *net.TCPConn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Recieved: %s\n\r", buf[:n])
		time.Sleep(time.Second)
	}
}

func main() {

	//var server_ip net.Addr = "10.100.23.147"

	//go UDP_listen(20016)
	//go UDP_send()
	go TCP_connect()

	select {}

}
