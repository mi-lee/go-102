Notes - July 26

- channels: specific type to go


## Types

`b1 := true `

- type inference done for you
- declare & assign vars

s1 := `Raw string literal`
- not escaped or backtick
s2 := "Interpreted string literal"

## Var declaration

make(T)         // make takes a type T, which must be a slice,
                // map or channel type, optionally followed by
                // a type-specific list of expressions

- slice: dynamic array (declare + preallocate memory)
- map: dictionary

### Nil

Not null. It is the value that means zero.

## Struct

Powerful types


type rectangle struct {
    width  int
    height int
}
- new type called rectangle, it is a struct with these fields
- can declare type functions, etc.
type <typename> <type description>

r1 := rectangle{1, 2}       // New rectangle with w + h


```
// dot notation and direct mutation
r1.width = 3                // Set width to a new value
fmt.Printf("Width = %d; Height = %d\n", r1.width, r1.height)

var r2 rectangle            // w=0, h=0 (int zero values)
r4 := rectangle{}           // w=0, h=0

// you should explicitly declare your params
r3 := rectangle{height: 1}  // w=0, h=1
```



## Variadics
- be careful - usually a smell
- means: invoke, you'll pass a comma separated set of strings

```
// explicitly returns int
func f5() int {             // Return type declaration
    return 42
}

```

```
<!-- very idiomatic -->
func f6() (int, string) {   // Multiple return values
    return 42, "foo"
}
```

- go does not have function overloading
- every function name is unique
- method is a function bound to a receiver

```
type rectangle struct {
    width  int
    height int
}

func (r rectangle) area() int {
    return r.width * r.height
}

func main() {
    r := rectangle{3, 4}
    fmt.Println(r.area())   // area 12
}
```


## Interfaces
- no implements keyword in go
- implementation is implicit





