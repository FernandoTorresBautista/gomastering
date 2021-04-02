package fanin

// takes multiples channels, abreviated mutation
func fain(chs ...<-chan int) <-chan int {
	out := make(chan int)
	for _, c := range chs {
		go registerChannel(c, out)
	}
	return out
}

// input channel c to the output channel out
func registerChannel(c <-chan int, out chan<- int) {
	for n := range c {
		out <- n
	}
}

// fan out channel,
// incomming input chanel on one channel and distribute the input message to multiples channels
//
// rooms, run method, where listenen every message and broadcast every message
// and distributed the message in every channel r.clients
