# Switch

Like `for` and `if`, you can run a short statement before the condition.

Go only runs the matching case - there is no fall-through - so there is no need for a break
statement.

```go
switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("Go runs on OS X.")
  case "linux":
    fmt.Println("Go runs on Linux.")
  default:
    fmt.Printf("Go runs on %s.", os)
}
```

Switch without a condition is a clean way to write if-elseif-else chains.

```go
switch {
  case now.Hour() < 12:
    fmt.Println("Good morning!")
  case now.Hour() < 17:
    fmt.Println("Good afternoon!")
  default:
    fmt.Println("Good evening!")
}
```
