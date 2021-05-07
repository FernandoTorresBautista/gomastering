package hydrarestapi

import (
	"GoMastering/Hydra/hydraconfigurator"
	"log"
	"net/http"
)

type DBlayerconfig struct {
	DB   string `json:"database"`
	Conn string `json:"connectionstring"`
}

func InitializeAPIHandlers() error {
	conf := new(DBlayerconfig)
	// get configuration factory function
	//err := hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, conf, "/../apiconfig.json")
	err := hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, conf, "./hydraweb/apiconfig.json")
	if err != nil {
		// log.Fatal crash the appication
		log.Println("Error decoding JSON", err)
		return err
	}
	h := newhydraCrewReqHandler()
	err = h.connect(conf.DB, conf.Conn)
	if err != nil {
		// log.Fatal crash the appication
		log.Println("Error connecting to db ", err)
		return err
	}
	http.HandleFunc("/hydracrew/", h.handleHydraCrewRequests)
	//http.Handle("/hydracrew/", h) -> enable the setverHTTP, less flexibility
	return nil
}

func RunAPI() error {
	if err := InitializeAPIHandlers(); err != nil {
		return err
	}
	return http.ListenAndServe(":8061", nil)
}
