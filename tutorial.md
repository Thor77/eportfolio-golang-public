Tutorial
========

This Tutorial will just cover some basic aspects of the language.
If you found this by accident you probably want to read the [official documentation](https://golang.org/doc/) instead.

# Basics
## Types
```go
bool // example: true
string // example: "string"

int  int8  int16  int32  int64 // example: 0
uint uint8 uint16 uint32 uint64 uintptr // example: 0

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point
     // example: 'a'

float32 float64 // example:  0.1

complex64 complex128 // example: (0+0i)
```

You can convert between types by using the type name as a function name:
```go
s := "abc"
sAsByteArray := []byte(s)
```

## Variables
You can define variables using an explicit notation by using `var` with a type and assigning a value later
```go
var x int
x = 1
```
define a value  directly
```go
var x = 1
```
or if you are inside a function omit the `var` completely
```go
x := 1
```

Constants are defined using
```go
const c1 = 0
```
but only chars, strings, booleans and integer can be constants.

## Functions
As for variables the type follows the argument name.
There are no optional arguments and no overloading so you have to use multiple functions with different names.
```go
func func1(arg1 int) int {
    return arg1
}
```
A function can return multiple values at once:
```go
func func2(arg1, arg2 int) (int, int, error) {
    return arg1, arg2, nil
}
```
and if there are multiple arguments with the same type, they can be combined.

# Flow Control
A `for`-loop looks like this:
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```
Any argument can be omitted:
```go
i := 0
for ; i < 10; {
    fmt.Println(i)
    i += 1
}
```
The above is equivalent to (so `for` is go's `while` as well)
```go
i := 0
for i < 10 {
    fmt.Println(i)
    i += 1
}
```
An endless loop can be defined without any argument
```go
for {
    // ever
}
```
An if looks like this:
```go
if i < 10 {

}
```
An if/else like this:
```go
if i < 10 {

} else {

}
```
`If`s can take a statement *and* a condition:
```go
if i:= 5; i < 3 {

}
```
There's also a `switch`-statement:
```go
abc := "abc"
switch abc {
case "abc":
    fmt.Println(":)")
case "cba":
    fmt.Println("uhm")
default:
    fmt.Println(abc)
}
```
which will exit after the first match.
It can be used as a replacement for if/else-trees:
```go
i := 0
switch {
case i < 10:
case i < 5:
case i < 1:
}
```
The `defer`-statement can be used to execute a statement at the end of a function:
```go
func main() {
    defer fmt.Println("Second")
    fmt.Println("First")
}
```
It follows last-in-first-out execution order.

# Advanced Types
A struct is a container for multiple values:
```go
type Point struct {
    X, Y int
}

func main() {
    p := Point{X: 1, Y: 2}
    // equivalent
    p = Point{1, 2}
    fmt.Println(p)
    fmt.Println(p.X, p.Y)
}
```
Arrays look like this and can only contain values of one type and have a static length:
```go
var a [2]int
a[0] = 0
a[1] = 1
// equivalent
a = [2]int{0, 1}
```
Slices are views on arrays
```go
a := [5]int{0, 1, 2, 3, 4}
sl := a[1:4]
fmt.Println(sl)
sl[0] = 42
fmt.Println(a)
```
Modifying them will also modify the underlying array.
They can be initialized with aa starting length and extended:
```go
sl := make([]int, 1)
fmt.Println(sl[0])
sl[0] = 1
sl = append(sl, 2)
fmt.Println(sl)
```
The `for`-loop can be used to iterate over a slice:
```go
sl := []string{"Hello", "World"}
for idx := range sl {
}
```
without the value or with the value
```go
for idx, value := range sl {
}
```
the index can be omitted
```go
for _, value := range sl {
}
```
There are `maps` which represent an association between a key and value:
```go
m := make(map[string]int)
m := map[string]int{
    "abc": 0,
}

// insert/update
m["def"] = 0
// retrieve
m["def"]
// delete
delete(m, "def")
// test key is present
value, isPresent := m["ghi"]
```
Pointers point to the address of an object
```go
i := 1
p := &i // pointer to i
fmt.Println(*p) // read i through p
*p = 2 // modify i through p
```
They are useful for passing them to functions to modify objects in place
```go
func Clear(p *Point) {
    p.X = 0
    p.Y = 0
}
```

# Methods
There are no classes, but methods with `receivers`:
```go
type Point struct {
    X, Y int
}

func (p Point) Sum() int {
    return p.X + p.Y
}

// fancy alternative for
func Sum(p Point) int {
    return p.X + p.Y
}

func main() {
    p := Point{1, 2}
    p.Sum()
    // fancy alternative for
    Sum(p)
}
```
`Pointer-receivers` are used to modify the object:
```go
func (p *Point) IncX() {
    p.X = p.X + 1
}

func (p *Point) IncY() {
    p.Y = p.Y + 1
}

func (p *Point) Clear() {
    p.X = 0
    p.Y = 0
}

func main() {
    p.IncX()
    p.IncY()
}
```

# Interfaces
Interfaces are a collection of methods
```go
type Clearable interface {
    Clear()
}
```
They can be used to assert that a given object implements a set of methods:
```go
func clearMe(c Clearable) {
    c.Clear()
}

func main() {
    clearMe(&Point{1, 2})
}
```
The `nil`-interface is implemented by all types
```go
interface{}
// dynamic typing \o/
fmt.Printf("%s, %d", "Hello World", 1)
```
it's only useful for a very limited set of usecases like `fmt.Printf` or JSON-decoding.
An empty interface can be converted to a type:
```go
var s interface{} = 1

sAsString := s.(string) // panic!
sAsInt := s.(int)
```
a `case` statement can be used to check for it's type:
```go
switch s.(type) {
case int:
case string
case bool
}
```

# Error handling
Errors need to be handled explicitly:
```go
_, err := strconv.Atoi("77")
if err != nil {
    log.Fatal(err)
}
```
There's an error package to simplify creating custom errors:
```go
func add(i1, i2 int) (int, error) {
    if r := i1 + i2; r < 100 {
        return r, nil
    } else {
        return 0, errors.New("oO")
    }
}

func main() {
    _, err := add(1, 99)
    if err != nil {
        log.Fatal(err)
    }
}
```

# Goroutines
To run code in parallel there are Goroutines, they are invoked with the `go` keyword and a function:
```go
func calculate(i int) {
    result := 0
    for ; i < 100; i++ {
        result = result + i + 42
    }
    fmt.Printf("I'm done: %d\n", result)
}

func main() {
    go calculate(5)
    fmt.Println("Done")
}
```
To synchronize Goroutines channels exist:
```go
func calculate(i int, results chan int) {
    result := 0
    for ; i < 100; i++ {
        result = result + i + 42
    }
    results <- result
}

func main() {
    results := make(chan int)
    go calculate(5, results)
    go calculate(10, results)
    result1, result2 := <- results, <- results
    fmt.Printf("Result1: %d Result2: %d\n", result1, result2)
}
```
A for statement can be used to read from a channel until it is closed:
```go
func sendSequential(s string, transport chan rune) {
    for _, val := range s {
        transport <- val
    }
    close(transport)
}

func main() {
    transport := make(chan rune)
    go sendSequential("Hello World", transport)
    for val := range transport {
        fmt.Print(string(val))
    }
    fmt.Print("\n")
}
```

# Packages
Every Go program is a package.
A package can consist of multiple files in the same folder,
The package name is not implied from the path but specified explicitly. The convention is to use the last part of the import path, though.
```go
package main

func main() {
}
```

The `main` package has a special meaning, it's main-function is the entrypoint for an application.

# Imports
Packages are imported using the `import` statement:
```go
import "fmt"
```
Packages from the stdlib can be executed like this, external packages are imported by their URL:
```go
import "github.com/gorilla/mux"
```

# Scopes
Scopes are defined by the capitalization of the objects:
```go
var a // unexported
var B // exported
type C struct { // exported

}
func func1() { // unexported

}
```

# Dependency Management
External packages are installed using `go get <path>`,
they will be downloaded into `$GOPATH/src`.

Because there's no simple way to specify a package version,
there is [`Go Modules`](https://github.com/golang/go/wiki/Modules) in recent releases.

It defines dependencies and their version in a `Go.mod` file.

# Tooling
There are multiple tools to simplify development.

`go fmt` can format code complying to Go's style guide.

`go vet` can detect and report various runtime errors.

`go imports` can automatically manage imports by adding/removing them as needed.
