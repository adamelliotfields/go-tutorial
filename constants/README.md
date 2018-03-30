# Constants

Constants can be of type `char`, `string`, `bool`, or `int`.

Typically, constants use CamelCase names. Since upper-case identifiers are exported, use lower
camelCase to make private constants.

```go
const Pi = 3.14
```

Numeric constants are high-precision values.

```go
const (
  Big = 1 << 100
  Small = Big >> 99
)
```
