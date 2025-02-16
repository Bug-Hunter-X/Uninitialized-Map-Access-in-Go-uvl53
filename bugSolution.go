```go
package main

import (


    "fmt"
)

func main() {
    var m map[string]int
    m = make(map[string]int)
    fmt.Println(m["a"]) //This will print 0
    if val, ok := m["a"]; ok {
        fmt.Println("Value of a:", val)
    } else {
        fmt.Println("Key a does not exist")
    }

    m["b"] = 10
    fmt.Println(m["b"])
    if val, ok := m["b"]; ok {
        fmt.Println("Value of b:", val)
    } else {
        fmt.Println("Key b does not exist")
    }

}
```