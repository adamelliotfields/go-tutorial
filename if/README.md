# If

If statements in Go are straight-forward. Expressions are not surrounded by parens, but braces are
required.

```go
if true {
  fmt.Println("true")
}
```

You can also execute a short statement before the condition. Variables declared inside the short
statement are available inside any `else` blocks.

```go
if c := true; c {
  fmt.Println("true")
}
```
