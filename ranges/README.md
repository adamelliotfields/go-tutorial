# Ranges

The `range` form of the `for` loop iterates over a slice or map.

`range` returns two values each iteration: the index and the value of the element at that index.

If you don't need to use the index, you can use `_` in its place.

```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

for i, v := range pow {
  fmt.Printf("2**%d = %d\n", i, v)
}
```
