# Interfaces

An interface type is a set of method signatures. Any type that defines all of the methods in the
interface will implicitly implement the interface. Unlike other object-oriented languages, there is
no "implements" keyword in Go.

A good naming convention for interfaces is to end in "-er", since interfaces describe actions. In
this example, anything that implements the JSONPrinter interfaces is a "JSON Printer".

```go
type JSONPrinter interface {
  toJSON() string
}
```

Instead of using a concrete type as a function parameter, use an interface type. Any type that
implements the interface can be passed as an argument.

```go
func printJSON(j JSONPrinter) {
  fmt.Println(j.toJSON())
}
```

Interface values can be thought of as a tuple of (value, type).

There also exists the empty interface type, which can be used as a wildcard type, like <?> in Java
or any in TypeScript. This works because every type in Go implements at least zero methods.

```go
// Prints the value and type of the underlying type implementing the interface.
func describe(i interface{}) {
  fmt.Printf("(%v, %T)\n", i, i)
}
```

Type assertions provide access to an interface's underlying concrete value. The type assertion
statement asserts that the interface value `i` holds the concrete type `string`.

Type assertions return two values: the underlying value and a boolean that reports whether the
assertion succeeded.

```go
var i interface{} = "hello world"

var s, ok = i.(string)

func main() {
  if ok {
    fmt.Printf("%s\n", s) // hello world
  }
}
```

A _type switch_ is a construct that permits several type assertions in series.

```go
func do(i interface{}) {
  switch v := i.(type) {
    case int:
      fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
      fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
      fmt.Printf("I don't know about type %T!\n", v)
  }
}

func main() {
  do(21) // Twice 21 is 42
  do("hello") // "hello" is 5 bytes long
  do(true) // I don't know about type bool!
}
```

One of the most ubiquitous interfaces is `Stringer`, defined byt the `fmt` package.

A `Stringer` is a type that can describe itself as a string. Many packages look for this interface
to print values. Any type that has a `String()` method implements the `Stringer` interface.

```go
type Person struct {
  Name string
  Age  int
}

func (p Person) String() string {
  return fmt.Sprintf("%v (%v years old)", p.Name, p.Age)
}

func main() {
  a := Person{"Adam Fields", 33}

  fmt.Printf("%s\n", a) // Adam Fields (33 years old)
}
```

Like the `Stringer` interface, the `error` interface has an `Error()` method that returns a string.

```go
type NegativeNumberError struct {
  i int
}

func (e *NegativeNumberError) Error() string {
  return fmt.Sprintf("Error: %d is a negative number", e.i)
}

func factorial(i int) (int, error) {
  var err error
  f := 1
  
  if i <= 0 {
    err = &NegativeNumberError{i}
    return 0, err
  }

  for i > 1 {
    f *= i
    i--
  }

  return f, err
}

func main() {
  if f, err := factorial(-3); err == nil {
    fmt.Println(f)
  } else {
    fmt.Println(err.Error()) // Error: -3 is a negative number
  }
}
```
