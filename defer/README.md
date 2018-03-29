# Defer

Defer is used to ensure that a function call is performed later in a program's execution.

Defer _defers_ the execution until the surrounding function returns.

```go
func deferredHelloWorld() {
  defer fmt.Println("World")

  fmt.Println("Hello")
}
```

Deferred calls are pushed onto a stack and executed in a last-in-first-out order.

```go
func deferredCounter() {
	fmt.Println("Counting...")

	for i := 1; i <= 10; i++ {
		defer fmt.Println(i) // 10 9 8 7 6 5 4 3 2 1
	}

	fmt.Println("Done!")
}
```
