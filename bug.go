```go
func main() {
    var m map[string]int
    fmt.Println(m["a"]) // This will not panic, but print 0
}
```