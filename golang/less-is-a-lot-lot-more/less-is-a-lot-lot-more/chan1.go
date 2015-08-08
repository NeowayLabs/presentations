// Declaring and initializing.
var c chan int
c = make(chan int)
// or
c := make(chan int)

// Sending on a channel.
c <- 1

// Receiving from a channel.
// The "arrow" indicates the direction of data flow.
value = <-c
