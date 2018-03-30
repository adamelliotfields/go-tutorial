# Methods

Go does not have classes, so there are no instance methods. However, you can define methods on
types.

A method is a function with a "receiver" parameter. The receiver appears between the `func` keyword
and the method name.

This is known as a "value" receiver.

```go
type Greeting struct {
  Message string
}

func (g Greeting) Greet() {
  fmt.Println(g.Message)
}
```

A "pointer reciever" acts like a setter method in traditional OOP. Pointer receivers are
recommended when possible, as they don't need to copy the value on each call.

It is also recommended to only use one type of receiver on a given type, because an interface cannot
define mixed types of receivers.

Note that even though the receiver is a pointer, you can still use a value. This is syntactic sugar
so you don't need to write `(&g).SetGreeting()`.

If a function takes a pointer as an argument, then you must pass a pointer.

```go
func (g *Greeting) SetGreeting(message string) {
  g.Message = message
}

func setGreeting(g *Greeting, message string) {
  g.Message = message
}

g.SetGreeting("hello world")

setGreeting(&g, "good bye world")
```

Methods can be declared on non-structs as well.

```go
type Name string

// Don't use a pointer here, as `n` will reference the memory address.
func (n Name) Print() {
  fmt.Println(n)
}
```
