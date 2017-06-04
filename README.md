# Golang Tools

I'm newcomer to go programming language. I'm try to learn go and write useful
packages.

## Crypto

A package for cryptographic needs in go programming language.

### 1. RC4

Example:

```go
package main

import (
    "fmt"

    "github.com/733amir/crypto"
)

func main() {
    // First create an object of RC4.
    cipher := crypto.RC4{}
    // Then initial it with a key from 1 to 256 bytes.
    cipher.Init([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8})
    // KeyStream return a function that each call of it will produce a byte to
    // XOR will your data and get encrypted data.
    next := cipher.KeyStream()
    fmt.Print("Keys are: ")
    for i := 0; i < 10; i++ {
        fmt.Printf("%x ", next())
    }
    fmt.Println()
}
```
