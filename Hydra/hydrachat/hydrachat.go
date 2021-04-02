package hydrachat

import (
	"GoMastering/Hydra/hlogger"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var logger = hlogger.GetInstance()

func Run(connection string) error {
	//l, err := net.Listen("tcp", ":2100")
	l, err := net.Listen("tcp", connection)
	if err != nil {
		logger.Println("Error connecting to chat client", err)
		return err
	}
	r := CreateRoom("HydraChat")
	//
	/*go func(l net.Listener) {
		for {
			conn, err := l.Accept()
			if err != nil {
				logger.Println("Error accepting connection from chat client", err)
				break
			}
			go handleConnection(r, conn)
		}
	}(l)*/

	// clean up the resources and signal the exiting
	go func() {
		// Handle SIGINT and SIGTERM.
		ch := make(chan os.Signal)                         //
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM) // interrup or end...
		<-ch

		l.Close()
		fmt.Println("Closing tcp connection")
		close(r.Quit)
		if r.ClCount() > 0 {
			<-r.Msgch
		}
		os.Exit(0)
	}()

	//
	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Println("Error acceptins connection from chat client: ", err)
		}
		go handleConnection(r, conn) //
	}

	return err
}

func handleConnection(r *room, c net.Conn) {
	logger.Println("Received request from client", c.RemoteAddr())
	r.AddClient(c)
}
