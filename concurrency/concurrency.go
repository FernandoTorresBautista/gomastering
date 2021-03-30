package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Section 3 > Video 12 Concurrency in go - gorutines
//	/Code_Files/Section_3_Code/../3.3-*
// Section 3 > Video 13 Concurrency in go - Channels
//  /Code_Files/Section_3_Code/../3.4-*
// Section 3 > Video 14 Concurrency in go - Select
//  /Code_Files/select.go
func main() {

	// test on channels
	//test01()
	// buffer channel
	//test02()
	//
	//test03()

	// select
	testSelect()
}

var scMapping = map[string]int{
	"James": 5,
	"Kevin": 10,
	"Rahul": 9,
}

func findSC(name, server string, c chan int) {
	//Simulate searching
	time.Sleep(time.Duration(rand.Intn(50)) * time.Minute)

	// return security clearance from map
	c <- scMapping[name]
}

func testSelect() {

	rand.Seed(time.Now().UnixNano())

	c1 := make(chan int)
	c2 := make(chan int)

	name := "James"

	go findSC(name, "Server 1", c1)
	go findSC(name, "Server 2", c2)

	select {
	case sc := <-c1:
		fmt.Println(name, " has a security clearance of ", sc, "found in server 1")
	case sc := <-c2:
		fmt.Println(name, " has a security clearance of ", sc, "found in server 2")
	//case <-time.After(10000 * time.Millisecond):
	//case <-time.After(1 * time.Minute):
	//	fmt.Println("Search timed out!!")
	default:
		fmt.Println("Too slow!!")
	}
}

func test03() {
	c := make(chan string)
	go sayHelloMultipleTimes(c, 5)
	for s := range c {
		fmt.Println(s)
	}

	v, ok := <-c
	fmt.Println("Channel closed? ", !ok, v)
}

func sayHelloMultipleTimes(c chan string, n int) {
	for i := 0; i <= n; i++ {
		c <- "Hello"
	}
	close(c)
}

func test02() {
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// range and close when use channels r important
}

func test01() {
	helloPrinted := make(chan bool)
	go waitAndSay(helloPrinted, "world")
	fmt.Println("Hello")

	helloPrinted <- true

	if b := <-helloPrinted; b {
		fmt.Println("Program now signalled to exit")
	}
}

func waitAndSay(c chan bool, s string) {
	if b := <-c; b {
		fmt.Println(s)
	}
	c <- true
}

//
