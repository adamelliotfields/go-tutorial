# Functions

A function can take zero or more arguments.

When defining a function, the parameter type comes after the variable name. When consecutive
parameters share a type, you only need to provide the type to the last one.

```go
func add(x, y int) int {
  return x + y
}
```

Functions can also be passed to other functions.

```go
func compute(fn func(x, y int) int) int {
  return fn(42, 42)
}
```

Functions can return multiple results.

```go
func swap(x, y string) (string, string) {
  return y, x
}

var y, x = swap(42, 42)
```

Return values can be "named". A return statement without arguments returns the named values.

```go
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x

  return
}
```

Go functions can be closures.

```go
func fibonacci(f1, f2 int) func() int {
  var first, second, next int

  // Fibonacci can be seeded with either 0,1 or 1,1
  first = f1
  second = f2

  return func() int {
    next = first
    first = second
    second = first + next

    return next
  }
}

func main() {
  f := fibonacci(0, 1)

  fmt.Println("Fibonacci:")

  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
```

Go also supports recursion, but it is NOT optimized for tail recursion. Generally speaking, Go is
not a functional programming language.
