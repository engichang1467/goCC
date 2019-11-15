# Variables

Here are the main types

```go
string
bool
int
int 	 int8	 int16	 int32	 int64
uint	 uint8	 uint16	 uint32  uint64  uintptr
byte     // alias for uint8
rune     // alias for int32
float32 float64
complex64 complex128
```

Finding the datatype for a variable
```go
import "fmt"

fmt.Printf("%T", varaible)
```

Stating variables

```go
var variable = value
// Optional:
var variable dataTypes = value

// Shorthand
variable := value

// 2 variables at the same line
variable1, variable2 := value1, value2
```