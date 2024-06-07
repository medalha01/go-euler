In Go, imports and packages are fundamental to organizing and modularizing code. Here’s an overview of how they work, along with instructions on how to run, build, and execute Go programs.

### Packages and Imports

#### Packages

1. **Standard Library**: Go comes with a rich standard library that includes packages for I/O, text processing, cryptography, and more.
2. **Custom Packages**: You can create your own packages to organize your code.

A Go package is a collection of source files in the same directory that are compiled together. A package is defined by its directory.

#### Importing Packages

To use functions from other packages, you need to import them. Import statements are used to bring external packages into your code.

```go
package main

import (
    "fmt"
    "math"
)
```

### Example: Creating and Using a Package

#### Step 1: Create a Package

Create a directory structure for your project. For example:

```
myproject/
    main.go
    mathutils/
        mathutils.go
```

In `mathutils/mathutils.go`:

```go
package mathutils

// Add two integers
func Add(a, b int) int {
    return a + b
}

// Subtract two integers
func Subtract(a, b int) int {
    return a - b
}
```

#### Step 2: Import and Use the Package

In `main.go`:

```go
package main

import (
    "fmt"
    "myproject/mathutils"
)

func main() {
    sum := mathutils.Add(3, 4)
    difference := mathutils.Subtract(10, 6)
    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", difference)
}
```

### Building and Running Go Programs

#### Running a Go Program

To run a Go program directly:

```sh
go run main.go
```

This compiles and runs the program in one step.

#### Building a Go Program

To build a Go program, which compiles the source code into an executable file:

```sh
go build main.go
```

This creates an executable file named `main` (or `main.exe` on Windows) in the current directory.

#### Running the Built Executable

After building the executable, you can run it directly from the command line:

```sh
./main
```

On Windows:

```sh
main.exe
```

### Managing Dependencies

Go modules (`go.mod` and `go.sum`) are used to manage dependencies in your project.

#### Step 1: Initialize a New Module

Navigate to your project directory and initialize a new module:

```sh
go mod init myproject
```

This creates a `go.mod` file which tracks your project's dependencies.

#### Step 2: Adding Dependencies

When you import a package that is not part of the standard library, you can add it to your project by running:

```sh
go get <package-path>
```

For example:

```sh
go get github.com/gorilla/mux
```

This updates your `go.mod` and `go.sum` files to include the new dependency.

### Full Example

Here’s a full example that includes importing standard packages, creating and using a custom package, and building/running the program.

#### Directory Structure

```
myproject/
    go.mod
    main.go
    mathutils/
        mathutils.go
```

#### `mathutils/mathutils.go`

```go
package mathutils

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}
```

#### `main.go`

```go
package main

import (
    "fmt"
    "myproject/mathutils"
)

func main() {
    sum := mathutils.Add(3, 4)
    difference := mathutils.Subtract(10, 6)
    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", difference)
}
```

#### Commands to Run

1. Initialize the module:

    ```sh
    go mod init myproject
    ```

2. Run the program:

    ```sh
    go run main.go
    ```

3. Build the program:

    ```sh
    go build main.go
    ```

4. Run the built executable:

    ```sh
    ./main
    ```

By following these steps, you can organize your Go code using packages, manage dependencies, and build and run your Go programs effectively.