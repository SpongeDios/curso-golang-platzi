package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingCLients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	host2 = flag.String("h", "localhost", "host")
	port2 = flag.Int("p", 3090, "port")
)

//Client1 -> Server -> HandleConnection(Client1)

func HandleConnection(conn net.Conn)  {
	defer conn.Close()
	message := make(chan string)
	go MessageWrite(conn, message)
	//clientName example: localhost:8080
	clientName := conn.RemoteAddr().String()
	message <- fmt.Sprintf("Welcome to the server, your name is: %s\n", clientName)
	messages <- fmt.Sprintf("New Client is here, name %s\n", clientName)
	incomingCLients <- message

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}

	leavingClients <- message
	messages <- fmt.Sprintf("%s dijo adios", clientName)
}

func MessageWrite(conn net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprintln(conn, message)
	}
}

func Broadcast(){
	clients := make(map[Client]bool)
	for {
		select {
		case message := <-messages:
			for client := range clients {
				client <- message
			}
		case newClient := <- incomingCLients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d",*host2, *port2))
	if err != nil {
		log.Fatal(err)
	}

	go Broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go HandleConnection(conn)
	}
}

