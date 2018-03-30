# Types

Go's basic types are:
  - `bool`
  - `string`
  - `int`, `int8`, `int16`, `int32`, `int64`
  - `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
  - `byte` _(alias for `uint8`)_
  - `rune` _(alias for `int32`)_
  - `float32`, `float64`
  - `complex64`, `complex128`

`int`, `uint`, and `uintptr` are usually 32-bits wide on 32-bit systems, and 64-bits wide on 64-bit
systems.

When you need an integer, it's recommended to just use `int` unless you have a specific reason for
using a sized or unsigned integer type.
