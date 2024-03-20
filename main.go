package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{} // 0 bytes
	msgch  chan string
}

func NewServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string),
	}
}

func (s *Server) start() {
	fmt.Printf("Server started\n")
	s.loop()
}

func (s *Server) loop() {
mainLoop:
	for {
		select {
		case <-s.quitch:
			break mainLoop
		case msg := <-s.msgch:
			s.handleMessage(msg)
		default:
		}
	}
	fmt.Println("server is shutting down gracefully ")
}

func (s *Server) handleMessage(msg string) {
	fmt.Printf("Message received: %s\n", msg)
}

func (s *Server) SendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) quit() {
	fmt.Println("quit command ")
	s.quitch <- struct{}{}
}

func main() {

	server := NewServer()

	go func() {

		fmt.Println("Sleep()")
		time.Sleep(5 * time.Second)
		server.quit()

	}()

	server.loop()
	fmt.Println("Programer  Orver")

}
