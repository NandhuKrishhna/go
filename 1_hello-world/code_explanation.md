This is a simple Go program that prints `"Learning go language"` to the console. Let's break it down.

---

## Code

```go
package main

import "fmt"

func main() {
    fmt.Println("Learning go language")
}
```

---

## Explanation

### `package main`

* Declares that this file is part of the `main` package.
* This tells Go that the program is executable and not a shared library.
* The `main` package is required for running the Go program.

### `import "fmt"`

* Imports the **fmt** package, which provides I/O formatting functions.
* Used here to access `fmt.Println` to print output to the console.

### `func main()`

* Defines the **main function**, the entry point of the Go program.
* When the program is run, execution begins from this function.

### `fmt.Println("Learning go language")`

* Calls the `Println` function from the `fmt` package.
* Prints `"Learning go language"` followed by a newline.

### `}`

* Closes the `main` function.

---

## Running the Code

### Step 1: Save the File

Save the code in a file with a `.go` extension, for example: `main.go`.

### Step 2: Run Using Terminal

Navigate to the directory containing the file and run:

```bash
go run main.go
```

This will:

* Compile the code.
* Execute the `main()` function.
* Output:

  ```
  Learning go language
  ```

---

## Additional Notes

### `go run`

* Compiles and runs the program in one step.
* Great for quick testing.

### `go build`

To create an executable file:

```bash
go build main.go
```

Then run the output:

* On Linux/macOS: `./main`
* On Windows: `main.exe`

---

## Summary

* `package main` identifies the program as executable.
* `func main()` is the entry point.
* Without them, Go will treat the file as a library, not an application.

---


