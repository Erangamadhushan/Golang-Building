># What is Go?

- Go is a cross-platform, open source programming language
- Go can be used to create high-performance applications
- Go is a fast, statically typed, compiled language known for its simplicity and efficiency
- Go was developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson in 2007
- Go's syntax is similar to C++

># What is Go Used For?
- Web development (server-side)
- Developing network-based programs
- Developing cross-platform enterprise applications
- Cloud-native development

> # Why Use Go?
- Go is fun and easy to learn
- Go has fast run time and compilation time
- Go supports concurrency
- Go has memory management
- Go works on different platforms (Windows, Mac, Linux, Raspberry Pi, etc.)

```bash 
go version #  To search go is available or not in local machine
```

```bash
go mod init example.com/hello 
# to create go.mod file like package.json file in MERN or JavaScript based fullstack application
```

> ## Sample code segment

```golang
// main.go file

package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}
```

To run the main.go fle

```bash
go run .\main.go
```

If you want to save the program as an executable, type and run:
```
go build ./main.go
```

> # Go Variable Types
## In Go, there are different types of variables, for example:

- `int`- stores integers (whole numbers), such as 123 or -123
- `float32`- stores floating point numbers, with decimals, such as 19.99 or -19.99
- `string` - stores text, such as "Hello World". String values are surrounded by double quotes
- `bool`- stores values with two states: true or false



> # Declaring (Creating) Variables

1. Use the var keyword, followed by variable name and type:

```golang
var variablename type = value
```
###  **Note: You always have to specify either type or value (or both).**


```golang
package main
import ("fmt")

func main() {
    var username string = "Eranga Madhushan"
    var userage int = 23

    fmt.Println("Current User Name : " + username + " and user age : " + fmt.Sprint(userage))
    
}

```

2. Use the `:=` sign, followed by the variable value:

```
variablename := value
```

**Note** : In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).

**Note** : It is not possible to declare a variable using :=, without assigning a value to it.

```golang
package main

import ("fmt")

func main() {
    inferradName := "Eranga Madhushan"
    inferradAge := 23

    fmt.Println(inferradName)
    fmt.Println(inferradAge)
}
```

