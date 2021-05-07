package main

import (
	"GoMastering/Hydra/hlogger"
	"GoMastering/Hydra/hydrachat"
	"GoMastering/Hydra/hydraweb/hydraportal"
	"flag"
	"strings"
	//"crypto/md5"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")

	// test only the portal
	//hydraportal.Run()

	// -o w => web server
	// -o c => tcp chat server
	operation := flag.String("o", "w", "Operation: w for web \n c for chat")
	flag.Parse()
	switch strings.ToLower(*operation) {
	case "c":
		err := hydrachat.Run(":2100")
		if err != nil {
			logger.Println("Could not run hydra chat", err)
		}
	case "w":
		err := hydraportal.Run()
		if err != nil {
			logger.Println("could not run hydra web portal", err)
		}
	}

}
