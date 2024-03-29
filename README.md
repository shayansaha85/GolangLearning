## go important commands
- **go build** : compiles bunch of go source codes, and generate executable file
- **go run** : compiles abd execute one or more codes
- **go fmt** : formats all the code in each file in the current directory
- **go install** : compiles and install package
- **go get** : downloads raw source code of 3rd party package
- **go test** : runs any test associated with current project

## break hello world code

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

```
- **package == project == workspace**
- every go file should have package name
- **why we call the package "main"?** : inside go there are 2 types of packages : **executable** (for coding) & **reusable** (used as helper. kind of dependencies)
- all files should have a main() function in order to run a code.
- **fmt** : it is part of standard library of go. link : golang.org/pkg
- **func** : it is a method name. short form for "function"

> **dtypes** : bool, string, int, float64