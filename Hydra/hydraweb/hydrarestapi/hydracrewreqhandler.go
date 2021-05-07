package hydrarestapi

import (
	"GoMastering/Hydra/hydradblayer"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type hydracrewreqhandler struct {
	dbConn hydradblayer.DBLayer
}

// constructor
func newhydraCrewReqHandler() *hydracrewreqhandler {
	return new(hydracrewreqhandler)
}

//
func (hcwreq *hydracrewreqhandler) connect(o, conn string) error {
	dblayer, err := hydradblayer.ConnectDatabase(o, conn)
	if err != nil {
		return err
	}
	hcwreq.dbConn = dblayer
	return nil
}

/*
func(hcwreq *hydracrewreqhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case "GET":
		ids := r.RequestURI[len("/hydracrew/"):]
		id, err := strconv.Atoi(ids)
		if err!= nil{
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w,"id %s provided is not of valid number. \n", ids)
			return
		}
		cm,err := hcwreq.dbConn.FindMember(id)
		if err!= nil{
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w,"Error %s occured when search for id %d \n ", err.Error(), id)
			return
		}
		json.NewEncoder(w).Encode(&cm)
	case "POST":
		cm := new(hydradblayer.CrewMember)
		err := json.NewDecoder(r.Body).Decode(cm)
		if err!= nil{
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w,"Error %s occured", err)
			return
		}
		hcwreq.dbConn.AddMember(cm)
	}
}
*/

func (hcwreq *hydracrewreqhandler) handleHydraCrewRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ids := r.RequestURI[len("/hydracrew/"):]
		id, err := strconv.Atoi(ids)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "id %s provided is not of valid number. \n", ids)
			return
		}
		cm, err := hcwreq.dbConn.FindMember(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error %s occured while searching for id %d \n ", err.Error(), id)
			return
		}
		json.NewEncoder(w).Encode(&cm)
	case "POST":
		cm := new(hydradblayer.CrewMember)
		err := json.NewDecoder(r.Body).Decode(cm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error %s occured", err)
			return
		}
		err = hcwreq.dbConn.AddMember(cm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error %s ocurred while adding a crew menber to the hydra database", err.Error())
			return
		}
		fmt.Fprintf(w, "Successfully inserted id %d \n", cm.ID)
	}
}
