# Arrays

An array is a numbered sequence of elements of a specific length. An array's length is part of its
type, so it cannot be resized.

```go
var words [2]string
```

You can also declare and initialize an array literal.

```go
var primes = [6]int{2, 3, 5, 7, 11, 13}
```

You can get and set element values using bracket notation.

```go
words[0] = "hello"

var two = primes[0]
```
