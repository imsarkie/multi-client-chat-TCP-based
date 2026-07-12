package server

import (
	"bufio"
	"fmt"
	"net"
)

type hub struct{
	clients 	map[net.Conn]bool
	uname 		map[string]net.Conn
	broadcast 	chan string
	register 	chan net.Conn
	unregister 	chan net.Conn 		// net.Conn is an interface
}

func newHub() *hub{
	return &hub{
		clients: 		make(map[net.Conn]bool),
		uname: 			make(map[string]net.Conn),
		broadcast: 		make(chan string),
		register: 		make(chan net.Conn),
		unregister: 	make(chan net.Conn),
	}
}

func handleUname(h *hub,scanner *bufio.Scanner, conn net.Conn) (string, bool){
	for{
		fmt.Fprintln(conn, "Please enter your Username: ")
		if !scanner.Scan(){
			return "", false // client disconnected before answering
		}
		username := scanner.Text()
		if _, exists := h.uname[username]; exists{
			fmt.Fprintln(conn, "Username already exists, Try new one.")
			continue
		} else {
			h.uname[username] = conn
			fmt.Fprintln(conn, "Yay! username is available...")
			return username, true
		}
	}
}

// No per-connection handler
func handleConn(h *hub, conn net.Conn){
	scanner := bufio.NewScanner(conn)

	username, ok := handleUname(h, scanner, conn)
	if !ok{
		return
	}

	h.register<- conn
	defer func ()  {
		h.unregister<- conn
		// here i need to delete the username if the connection is unregister
		delete(h.uname, username)
	}()

	for scanner.Scan(){
		msg := scanner.Text()
		if msg == "QUIT"{
			return
		}
		h.broadcast<- fmt.Sprintf("[%s]: %s", username, msg)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error: ", err)
	}
}

func (h *hub) selectLoop(){
	for{
		select{
		case conn := <-h.register:
			h.clients[conn] = true
		case conn := <-h.unregister:
			delete(h.clients, conn)
			conn.Close()
		case msg := <-h.broadcast:
			for conn := range h.clients{
				fmt.Fprintln(conn, msg)
			}
		}
	}
}

func Run(){
	h := newHub()
	go h.selectLoop()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Listening for clients from port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error connecting to client: ", err)
			continue
		}
		go handleConn(h, conn)
	}
}
