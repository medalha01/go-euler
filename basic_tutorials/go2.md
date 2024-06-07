To import and use functions or variables from other files in Go, you need to organize your code into packages and then import those packages where needed. Here’s a detailed guide on how to achieve this.

### Step-by-Step Guide

#### Step 1: Organize Your Project Directory

Let's create a simple project structure where we have multiple Go files in different packages.

```
myproject/
    main.go
    utils/
        stringutils.go
        mathutils.go
```

#### Step 2: Create Packages in Separate Files

1. **stringutils.go**:

```go
package utils

// Reverse returns the reverse of the input string
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], j = runes[j], runes[i]
    }
    return string(runes)
}
```

2. **mathutils.go**:

```go
package utils

// Add returns the sum of two integers
func Add(a, b int) int {
    return a + b
}

// Subtract returns the difference of two integers
func Subtract(a, b int) int {
    return a - b
}
```

#### Step 3: Import and Use the Packages in main.go

```go
package main

import (
    "fmt"
    "myproject/utils" // Import the utils package
)

func main() {
    // Use functions from stringutils.go
    reversedString := utils.Reverse("Hello, World!")
    fmt.Println("Reversed String:", reversedString)

    // Use functions from mathutils.go
    sum := utils.Add(3, 4)
    difference := utils.Subtract(10, 6)
    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", difference)
}
```

### Step 4: Initialize a Go Module

Navigate to your project directory and initialize a new module. This step is crucial for managing imports and dependencies properly.

```sh
cd myproject
go mod init myproject
```

### Step 5: Run and Build the Project

1. **Run the Project**:

```sh
go run main.go
```

2. **Build the Project**:

```sh
go build
```

This will create an executable file (`myproject` or `myproject.exe` on Windows) in the `myproject` directory.

3. **Run the Built Executable**:

```sh
./myproject
```

### Explanation

1. **Package Declaration**: Each Go file starts with a `package` declaration. Files in the `utils` directory declare `package utils`, meaning they belong to the `utils` package.

2. **Import Statements**: In `main.go`, the `utils` package is imported using `import "myproject/utils"`. Go's import path is based on the module name and the directory structure.

3. **Function Usage**: Functions from the `utils` package are called using `utils.FunctionName`.

### Summary

To import and use functions from other files in Go:

1. Organize your code into directories representing different packages.
2. Ensure each file has the appropriate `package` declaration.
3. Import the necessary packages using the import path based on the module name and directory structure.
4. Use the functions or variables from the imported packages in your code.

This structure helps maintain clean and modular code, making it easier to manage and scale your Go projects.