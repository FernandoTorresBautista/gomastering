package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var errCrewNotFound = errors.New("crew member not found")

var scMapping = map[string]int{
	"James": 5,
	"Kevin": 10,
	"Rahul": 9,
}

type findError struct {
	Name, Server, Msg string
}

func (e findError) Error() string {
	return e.Msg
}

func findSC(name, server string) (int, error) {
	//Simulate searching
	//time.Sleep(time.Duration(rand.Intn(50)) * time.Minute)
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	if v, ok := scMapping[name]; !ok {
		return -1, findError{name, server, "Crew member not found"}
		//return -1, errors.New("Crew member not found")
		//return -1, fmt.Errorf("Crew member %s not found on server '%s'", name, server)
		//return -1, errCrewNotFound
	} else {
		return v, nil
	}

}

func main() {
	//
	//test01()
	//
	test02()
}

func test01() {
	rand.Seed(time.Now().UnixNano())
	clearance, err := findSC("Ruko", "Server 1")

	if err == errCrewNotFound {
		fmt.Println("Error crew not found ", errCrewNotFound)
	}

	fmt.Println("Clearance level found: ", clearance, ", Error code: ", err)
}

func test02() {
	rand.Seed(time.Now().UnixNano())
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A Panic recovered ", err)
		}
	}()

	if i, err := findSC("Ruku", "server 1"); err != nil {
		fmt.Println("Error occured while searching ", err)

		if err == errCrewNotFound {
			//do some logic here to handle the error
			fmt.Println("Confirmed error is errCrewNotFound")
		}

		if v, ok := err.(findError); ok {
			fmt.Println("Server name ", v.Server)
			fmt.Println("Crew member name ", v.Name)
		}
	} else {
		fmt.Println("Crew member has security clearance", i)
	}
}
