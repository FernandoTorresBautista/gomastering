package hydrachat

import (
	"bufio"
	"io"
)

// chat client object
type client struct {
	*bufio.Reader
	*bufio.Writer
	wc chan string // witch client is 
}

// pipeline pattern stage 1: 
// whe we start a new client
//		cuando sea que nosotros mandemos un mensaje recibido por un cliente al chat room
// 

//func StartClient(msgCh chan<- string, cn net.Conn, quit chan struct{}) (chan<- string, <-chan struct{}) {
func StartClient(msgCh chan<- string, cn io.ReadWriteCloser, quit chan struct{}) (chan<- string, <-chan struct{}) { // TCP connection 
	c := new(client)
	c.Reader = bufio.NewReader(cn) // the client connection 
	c.Writer = bufio.NewWriter(cn)
	c.wc = make(chan string)
	done := make(chan struct{})

	//setup the reader
	go func() {
		scanner := bufio.NewScanner(c.Reader)
		for scanner.Scan() {
			logger.Println(scanner.Text())
			msgCh <- scanner.Text()
		}
		done <- struct{}{}
	}()

	//setup the writer
	c.writeMonitor()

	go func() {
		select {
		case <-quit:
			cn.Close()
		case <-done:
		}
	}()

	return c.wc, done
}

// 
// 
func (c *client) writeMonitor() {
	go func() {
		for s := range c.wc {
			//logger.Println("Sending",s)
			c.WriteString(s + "\n")
			c.Flush()
		}
	}()
}
