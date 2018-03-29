# Pointers

Pointers are references to memory addresses where values are stored.

Modifying the pointer value will modify the original variable.

The `&` operator generates a pointer to its operand. It is also known as the _reference operator_.

```go
type Person struct{
  Name string
  Age int
}

var p = Person{"Adam", 0}

var q = &p

q.Age = 34

fmt.Printf("Age: %d\n", p.Age) // 34
```

A pointer is represented by the `*` operator. It is also known as the _dereference operator_.

You'll use this when creating "pointer receivers" for methods, and functions that only take pointers.

```go
func (p *Person) SetAge(age int) {
  p.Age = age
}

p.SetAge(34)

func setAge(p *Person, age int) {
  p.Age = age
}

setAge(&p, 34)
```
