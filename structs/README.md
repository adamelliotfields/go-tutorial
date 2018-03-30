# Structs

A `struct` is a collection of fields.

```go
type Person struct {
  Name string
  Age  int
}

var adam = Person{
  Name: "Adam",
  Age:  33
}
```

If you know the order of the fields, you don't have to declare them.

```go
var adam = Person{"Adam", 33}
```

If you only need a `struct` once, you don't have to create a separate type.

```go
var adam = struct{Name string; Age int}{"Adam", 33}
```

`struct` fields are accessed using dot notation or through a pointer.

```go
var name = adam.Name

var p = &adam

var age = p.Age
```

`struct` fields can also be functions.

```go
type Counter struct {
  increment func() int
  decrement func() int
}

func counter() Counter {
  v := 0

  return Counter{
    increment: func() int { v += 1; return v },
    decrement: func() int { v -= 1; return v },
  }
}

var c = counter()

c.increment() // 1
c.decrement() // 0
```
