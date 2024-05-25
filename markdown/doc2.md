# Getting Started with Go

To get started with Go, follow these steps:

1. Install Go on your system by downloading it from the official Go website (https://golang.org).

2. Set up your Go workspace. Create a directory for your Go projects and set the GOPATH environment variable to point to that directory.

3. Create a new Go file with a `.go` extension and write your Go code.

4. Use the `go run` command to run your Go program, or use `go build` to compile it into an executable.

Here's a simple "Hello, World!" program in Go:

```go
package main

import "fmt"

func main() {
   fmt.Println("Hello, World!")
}
