# Packages

Packages are how Go organizes source code into directories. The name of the package is the directory
name.

When developing an executable program, use the `package main`.

```go
package main
```

The `import` keyword is used to import packages.

Imports can be grouped into a "factored" import statement.

```go
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
```

If the first letter of an identifier is capitalized, it will be exported for use in other packages.

```go
var Pi = math.Pi
```
