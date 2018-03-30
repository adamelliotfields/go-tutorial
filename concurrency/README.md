# Concurrency

A _goroutine_ is a lightweight thread managed by the Go runtime.

The `go` keyword creates a new goroutine. 

```go
func print(s string) {
	fmt.Println(s)
}

func main() {
	go print("World")
  print("Hello")
  // Need to wait a small amount of time for the goroutine to return,
  //   otherwise, the process will exit before it can print.
	time.Sleep(time.Millisecond * 100)
}
```

Channels are a way to send and receive values between goroutines.

Each channel has a type associated with it. Channels must be created using `make()` before they can
be used.

Data is either sent to or received from a channel using the `<-` operator.

Sending and receiving is blocking. In the following example, `go print("World", ch)` is blocked
until the main goroutine receives data back from the `ch` channel.

```go
func print(s string, ch chan bool) {
	fmt.Println(s)
	ch <- true
}

func main() {
	ch := make(chan bool)

	go print("Hello", ch)
	<-ch
	go print("World", ch)
	<-ch
}
```
