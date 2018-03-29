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
