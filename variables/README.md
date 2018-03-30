# Variables

The `var` statement declares a list of variables. Like with function parameters, the type is last.

`var` can be used at the package or function level.

```go
var b bool

var i, n int
```

If a variable is initialized, the type can be omitted.

```go
var a, b = 1, 2
```

Variables can be initialized inside a function block.

```go
var x, y string

func main() {
  x = "hello"
  y = "world"

  fmt.Printf("%s %s\n", x, y)
}
```

Inside a function, you can use the short assignment statement to assign and initialize a variable.

```go
func main() {
  z := true

  if z {
    fmt.Println("true")
  }
}
```
