package main

import (
	"GoMastering/Hydra/hydradblayer/passwordvault"
	"crypto/md5"
	"fmt"
)

func main() {

	//
	db, err := passwordvault.ConnectPasswordVault()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	//
	ferpss := md5.Sum([]byte("ferpas"))
	minapss := md5.Sum([]byte("minaspass"))
	jimpass := md5.Sum([]byte("jimspass"))
	caropass := md5.Sum([]byte("carospass"))
	passwordvault.AddBytesToVault(db, "Fer", ferpss[:])
	passwordvault.AddBytesToVault(db, "Mina", minapss[:])
	passwordvault.AddBytesToVault(db, "Jim", jimpass[:])
	passwordvault.AddBytesToVault(db, "Caro", caropass[:])
	db.Close()

}
