package main

import (
	"GoMastering/Hydra/hydracommlayer"
	"GoMastering/Hydra/hydracommlayer/hydraproto"
	"fmt"

	"flag"
	"log"
	"strings"
)

func main() {
	op := flag.String("type", "", "Server (s) or client (c) ?")
	address := flag.String("addr", ":8080", "address? host:port")
	flag.Parse()

	fmt.Println("switch: ", *op, *address)
	switch strings.ToUpper(*op) {
	case "S":
		fmt.Println("server...")
		runServer(*address)
	case "C":
		fmt.Println("client...")
		runClient(*address)
	}
	fmt.Println("_")
}

func runServer(dest string) {
	fmt.Println("run server...")
	c := hydracommlayer.NewConnection(hydracommlayer.Protobuf)
	recvChan, err := c.ListenAndDecode(dest)
	if err != nil {
		log.Fatal(err)
	}
	for msg := range recvChan {
		log.Println("Received: ", msg)
	}
	log.Fatal("err: ", err)
}

func runClient(dest string) {
	c := hydracommlayer.NewConnection(hydracommlayer.Protobuf)
	ship := &hydraproto.Ship{
		Shipname:    "Hydra",
		CaptainName: "Anng",
		Crew: []*hydraproto.Ship_CrewMember{
			&hydraproto.Ship_CrewMember{1, "ken", 5, "Pilot"},
			&hydraproto.Ship_CrewMember{2, "Jade", 4, "Tech"},
			&hydraproto.Ship_CrewMember{2, "Wally", 3, "Enginner"},
		},
	}

	if err := c.EncodeAndSend(ship, dest); err != nil {
		log.Println("Error ocurred while sending message: ", err)
	} else {
		log.Println("Send operation successful")
	}
}

/*
	ship := &hydraproto.Ship{
		Shipname:    "Hydra",
		CaptainName: "Jala",
		Crew: []*hydraproto.Ship_CrewMember{
			&hydraproto.Ship_CrewMember{1, "Kevin", 5, "Pilot"},
			&hydraproto.Ship_CrewMember{2, "Jade", 4, "Tech"},
			&hydraproto.Ship_CrewMember{3, "Wally", 3, "Enginneer"},
		},
	}
	ship := &hydraThrift.Ship{
		Shipname:    "Hydra",
		CaptainName: "Jala",
		Crew: []*hydraThrift.CrewMember{
			&hydraThrift.CrewMember{1, "Kevin", 5, "Pilot"},
			&hydraThrift.CrewMember{2, "Jade", 4, "Tech"},
			&hydraThrift.CrewMember{3, "Wally", 3, "Enginneer"},
		},
	}
*/
