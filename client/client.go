package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func receive(conn net.Conn){
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error in scanning in client: ", err)
	}
}



func main(){
	fmt.Println("Connecting with server...")

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")
	// fmt.Println("Local Address :", conn.LocalAddr())
	fmt.Println("Remote Address:", conn.RemoteAddr())

	go receive(conn)
	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		msg := scanner.Text()
		fmt.Fprintln(conn, msg)
		if msg == "QUIT"{
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning in client: ", err)
	}

	defer conn.Close()
}