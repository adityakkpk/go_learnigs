# go_learnigs
This repo contains all the code files which I have write while learning GO Lang

## Day-1
- Setting UP the Go environment and Installation
- To Run Go program in any of the folder outside we need to initialize `Go module` first using command:
```bash
go mod init <module_name>
```
- Some more commands: 
    1. For dependencies download 
    2. For all the packages/dependencies download
    3. Build a file (Cross Platform Executable Generator)
    3. Build a folder having main.go for a specific OS
```bash
go get <package_name>

go mod download

go build <file_name>

GOOS=windows GOARCH=amd64 go build .
```

- Some conventions about package names:
    - Package name should be in small case
    - Package names have to be as short as possible

- Functions (some of the example):
    
```
func add(x int, y int) int {
	return x + y
}

func add(x, y int) int {
    return x + y
}
```

- Next i am following this documentation : [Go](https://go.dev/tour)


- There are 5 reference types are there in the Go:
    1. Pointers
    2. Slices
    3. Functions
    4. Channels
    5. Maps

- There are Two type of Aggregators (Aggregate DataTypes) in GO:
    1. Arrays (Similar types of Data)
    2. Structs (Different types of Data)

### Struct
```go
// Syntax
// type struct_name struct {
// variable_name dataType
// variable_name dataType ....
// }

// Example
import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
}
```
- The first letter should always be capital of all teh variables declared in the Struct.


### Arrays
```go
// Example
import "fmt"

func main() {
	var a [2]string
    a[0] = "Hello"
    a[1] = "World"

    fmt.Println(a[0], a[1])
    fmt.Println(a)
}
```
- Array's size is a part of its type. So arrays cannot be resized


### Slices 
- A slice is:
    - A dynamic, flexible view over an array
    - Does not store data itself
    - Points to an underlying array

```go
// Example
s := []int{1, 2, 3}

/*
This creates:
    - An array in memory
    - A slice header pointing to it
*/



// Slice Declaration using 'make'
s := make([]int, 3)        // len=3, cap=3 → [0 0 0]
s := make([]int, 3, 5)     // len=3, cap=5

len(s) // number of usable elements
cap(s) // total space from start pointer



s := make([]int, 3, 5)
// len = 3 → usable
// cap = 5 → total allocated



// A slice is a small struct:
type slice struct {
    ptr *T   // pointer to array
    len int
    cap int
}



// Slicing a slice (Shared Memory ⚠️)
s1 := []int{1, 2, 3, 4, 5} // len=5, cap=5
s2 := s1[1:4] // [2 3 4]

// ⚠️ Both share the same array
s2[0] = 99
fmt.Println(s1) // [1 99 3 4 5]



// How slices are dynamic:
s := []int{1, 2, 3} // len=3 cap=3
s = append(s, 4)

// Old Array: [1 2 3]
// New Array: [1 2 3 4 _ _] //After appending a element is doubled its size/capacity to (6) and length become (4)

```