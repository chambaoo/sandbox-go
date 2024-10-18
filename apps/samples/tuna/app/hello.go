package main

import (
    "fmt"
    "example/lib"  // This now works since lib is a separate Go module
)

func main() {
    fmt.Println("Hello from the main application!")
    lib.Menu()  // Call the Menu function from the lib package
}
