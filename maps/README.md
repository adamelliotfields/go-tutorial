# Maps

A map _maps_ keys to values.

`make` retuns an initialized map ready for use.

```go
var m = make(map[string]string)
```

You can also make a map literal.

```go
var m = map[string]string{
  "message": "hello world"
}
```

Use `map[key]` to insert or update an element in a map. You can also retrieve an element this way.

```go
var m = make(map[string]string)

m["message"] = "hello world"

fmt.Println(m["message"])
```

You can delete an element using `delete(map, key)`.

```go
delete(m, "message")
```

You can test that an element exists with a two-value assignment.

The first value is the element, the second is a boolean.

```go
if _, ok := m["message"]; ok {
  fmt.Println(ok)
}
```
