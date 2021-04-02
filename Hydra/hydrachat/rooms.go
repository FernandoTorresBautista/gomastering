package hydrachat

import (
	"fmt"
	"io"
	"sync"
)

type room struct {
	name    string
	Msgch   chan string
	clients map[chan<- string]struct{}
	Quit    chan struct{}
	*sync.RWMutex
}

func CreateRoom(name string) *room {
	r := &room{
		name:    name,
		Msgch:   make(chan string),
		RWMutex: new(sync.RWMutex),
		clients: make(map[chan<- string]struct{}),
		Quit:    make(chan struct{}),
	}
	r.Run()
	return r
}

//func (r *room) AddClient(c net.Conn) {
func (r *room) AddClient(c io.ReadWriteCloser) { // TCP connection 
	logger.Println("Adding client" ) //, c.RemoteAddr())
	r.Lock()
	wc, done := StartClient(r.Msgch, c, r.Quit)
	r.clients[wc] = struct{}{}
	r.Unlock()

	//remove client when done is signalled
	go func() {
		<-done
		r.RemoveClient(wc)
	}()
}

func (r *room) ClCount() int {
	return len(r.clients)
}

func (r *room) RemoveClient(wc chan<- string) {
	logger.Println("Removing client ")
	r.Lock()
	close(wc) // close the channel of the client
	delete(r.clients, wc) // delete it from the clients 
	r.Unlock()
	select {
	case <-r.Quit:
		if len(r.clients) == 0 { // if there is no more clients close the room chatroom
			close(r.Msgch)
		}
	default:
	}
}

// second stage pipeline
// cuando recibimos un mensaje en el canal de mensajes que pertenece 
// a nuestro objeto chatroom
// cuando mostramos a todos los clientes conectados el mensaje  

func (r *room) Run() {
	logger.Println("Starting chat room ", r.name)
	go func() { //
		for msg := range r.Msgch { // listenen any message from the channel
			r.broadcastMsg(msg) // brodcast to every client in the channel 
		}
	}() 
}

func (r *room) broadcastMsg(msg string) {
	r.RLock()
	defer r.RUnlock()
	fmt.Println("Received message: ", msg)
	for wc, _ := range r.clients {
		go func(wc chan<- string) {
			wc <- msg
		}(wc)
	}
}
