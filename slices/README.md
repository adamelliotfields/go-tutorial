# Slices

Unlike arrays, slices can be dynamically sized. A slice actually doesn't store any data - it just
describes a section of an underlying array.

```go
var a = [4]string{"John", "Paul", "George", "Ringo"}

var s = a[0:1] // [John]
```

You can also use slice defaults to omit the low or high bounds of its underlying array.

```go
var s = a[:1] // [John]
```

Changing the elements of a slice modifies the elements of its underlying array.

```go
s[0] = "Mick"

fmt.Println(a[0]) // "Mick"
```

A slice literal is like an array literal without the length. Slice literals create the underlying
array automatically and then build a slice of that array.

```go
var sl = []struct{Key int; Value bool}{{0, false},{1, true}}
```

You can also slice a slice.

```go
var sls = s[:1] // {0, false}
```

Slices have a length and a capacity. The length of a slice is the number of elements it contains.
The capacity is the number of elements in the underlying array, counting from the low-bound of the
slice.

```go
len(s) // 1

cap(s) // 4
```

The zero value of a slice is `nil`. A `nil` slice has a length and capacity of `0`.

```go
var n []int

len(n) // 0

cap(n) // 0
```

You can also use `make()` to create a slice with a specified type, length, and capacity.

`make()` creates a zeroed-array and then a slice of that array with the specified length.

```go
var m = make([]int, 5, 7) // [0 0 0 0 0]
```

You can append items to a slice using `append()`.

```go
n = append(n, 1, 2, 3) // [1 2 3]
```
