# Loops

Go only has one looping construct - the `for` loop. The basic `for` loop has three components:
  - Init: executed before the first iteration.
  - Condition: executed before every iteration.
  - Post: executed at the end of every iteration.

The loop ends when the condition evaluates to false.

```go
for i := 1; i <= 30; i++ {
  var fb = ""

  if i % 3 == 0 { fb += "Fizz" }
  if i % 5 == 0 { fb += "Buzz" }

  if fb != "" {
    fmt.Println(fb)

    // `continue` exits the loop and starts the next iteration.
    // `break` exits the loop, but continues function execution.
    // `return` exits the loop, and stops functions execution.
    continue
  }

  fmt.Println(i)
}
```

Go has no formal `while` loop, but a `for` loop can be used as such.

```go
var sum = 1

for sum < 1000 {
  sum += sum
}
```

A `for` loop without any condition will loop forever.

```go
var i = 1

for {
  time.Sleep(time.Second)
  fmt.Println(i)
  i += 1
}
```
