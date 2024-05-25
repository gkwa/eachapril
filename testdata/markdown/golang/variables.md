# Variables in Golang

In Golang, variables are declared using the `var` keyword followed by the variable name, type, and an optional initial value. The type of the variable determines the kind of data it can store.

Examples:
```go
var age int = 25
var name string = "John"
var isStudent bool = true
Golang also supports short variable declarations using the := operator, where the type is inferred from the initial value:
goCopy codecount := 10
message := "Hello, World!"
Variables can be declared as constants using the const keyword, which means their values cannot be changed once assigned:
goCopy codeconst PI float64 = 3.14159
const USERNAME string = "admin"
```

It's important to choose meaningful names for variables and follow the Golang naming conventions to make the code more readable and maintainable.