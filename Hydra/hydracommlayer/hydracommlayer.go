package hydracommlayer

import (
	"GoMastering/Hydra/hydracommlayer/hydraproto"
	"fmt"
)

// Communications messages types
const (
	Protobuf uint8 = iota
)

type HydraConnection interface {
	EncodeAndSend(obj interface{}, destination string) error
	ListenAndDecode(listenaddress string) (chan interface{}, error)
}

func NewConnection(connType uint8) HydraConnection {
	fmt.Println("new connection ... : ", connType)
	switch connType {
	case Protobuf:
		fmt.Println("protobuf ...")
		return hydraproto.NewProtoHandler()
	}
	fmt.Println("nil ...")
	return nil
}
