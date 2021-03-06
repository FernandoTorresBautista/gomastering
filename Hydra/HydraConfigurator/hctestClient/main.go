package main

import (
	"GoMastering/Hydra/hydraconfigutor"
	"fmt"
)

type ConfS struct {
	TS      string  `name:"testString" xml:"testString" json:"testString"`
	TB      bool    `name:"testBool"  xml:"testBool" json:"testBool"`
	TF      float64 `name:"testFloat"  xml:"testFloat" json:"testFloat"`
	TestInt int
}

// go run .\main.go - run from this folder

func main() {
	configstruct := new(ConfS)
	//hydraconfigutor.GetConfiguration(hydraconfigutor.CUSTOM, configstruct, "configfile.conf")
	//hydraconfigutor.GetConfiguration(hydraconfigutor.JSON, configstruct, "configfile.json")
	hydraconfigutor.GetConfiguration(hydraconfigutor.XML, configstruct, "configfile.xml")
	fmt.Println(*configstruct)

	if configstruct.TB {
		fmt.Println("bool is true")
	}

	fmt.Println(float64(4.8 * configstruct.TF))

	fmt.Println(5 * configstruct.TestInt)

	fmt.Println(configstruct.TS)
}
