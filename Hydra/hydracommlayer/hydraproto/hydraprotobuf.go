package hydraproto

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
)

type ProtoHandler struct{}

// constructor for the protocol buffer sender
func NewProtoHandler() *ProtoHandler {
	return new(ProtoHandler)
}

func (pSender *ProtoHandler) EncodeAndSend(obj interface{}, destination string) error {
	v, ok := obj.(*Ship)
	if !ok {
		return errors.New("proto: unknown message type")
	}
	data, err := proto.Marshal(v)
	if err != nil {
		return err
	}
	return sendmessage(data, destination)
}

func (pSender *ProtoHandler) DecodeProto(buffer []byte) (*Ship, error) {
	pb := new(Ship)
	return pb, proto.Unmarshal(buffer, pb)
}

func (pSender *ProtoHandler) ListenAndDecode(listenaddress string) (chan interface{}, error) {
	outChan := make(chan interface{})
	l, err := net.Listen("tcp", listenaddress)
	if err != nil {
		fmt.Println("Err: ", err)
		return outChan, err
	}
	log.Println("Listening to ", listenaddress)
	// go routine first the listener
	go func() {
		defer l.Close()
		fmt.Println("listener...")
		for {
			c, err := l.Accept()
			if err != nil {
				fmt.Println("error: ", err)
				break
			}
			log.Println("Accepted connection from ", c.RemoteAddr())
			// new go routine, handle the connection
			go func(c net.Conn) {
				defer c.Close()
				fmt.Println("connection...")
				for {
					buffer, err := ioutil.ReadAll(c)
					if err != nil {
						break
					}
					if len(buffer) == 0 {
						continue
					}
					obj, err := pSender.DecodeProto(buffer)
					if err != nil {
						continue
					}
					select {
					case outChan <- obj:
					//case <-time.After(1 * time.Second):
					default:
					}
				}
			}(c)
		}
	}()
	return outChan, nil
}

func sendmessage(buffer []byte, destination string) error {
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Printf("Sending %d butes to %s \n", len(buffer), destination)
	_, err = conn.Write(buffer)
	return err
}

/*func EncodeProto(obj interface{}) ([]byte, error) {
	if v, ok := obj.(*Ship); ok {
		return proto.Marshal(v)
	}
	return nil, errors.New("Proto: Unknown message type")
}

func DecodeProto(buffer []byte) (*Ship, error) {
	pb := new(Ship)
	return pb, proto.Unmarshal(buffer, pb)
}
*/
