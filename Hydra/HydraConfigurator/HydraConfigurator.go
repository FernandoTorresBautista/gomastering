package HydraConfigurator

// entry point for the package...
//

import (
	"errors"
	"reflect"
)

const (
	CUSTOM uint8 = iota
	JSON
	XML
)

// global error
var wrongTypeError error = errors.New("Type must be a pointer to a struct")

func GetConfiguration(confType uint8, obj interface{}, filename string) (err error) {
	//check if this is type pointer
	mysRValue := reflect.ValueOf(obj)
	if mysRValue.Kind() != reflect.Ptr || mysRValue.IsNil() { // if is not a pointer or is null is not valid entry
		return wrongTypeError
	}
	//get and confirm the struct value
	mysRValue = mysRValue.Elem()
	if mysRValue.Kind() != reflect.Struct {
		return wrongTypeError
	}

	switch confType {
	case CUSTOM:
		err = MarshalCustomConfig(mysRValue, filename)
	case JSON:
		err = decodeJSONConfig(obj, filename)
	case XML:
		err = decodeXMLConfig(obj, filename)
	}
	return err
}
