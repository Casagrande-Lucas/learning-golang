# 11. Common Pitfalls and How to Avoid Them

### 1. **Unintended Variable Shadowing**

```
func example() {
    var name string = "Alice"
    if true {
        name := "Bob" // Shadows the outer 'name'
        fmt.Println(name)
    }
    fmt.Println(name) // Outputs: Alice
}
```
**Solution:** Use distinct variable names or avoid redeclaring variables in inner scopes.

### 2. **Nil Pointer Dereferencing**

```
var p *int
fmt.Println(*p) // Runtime panic: invalid memory address or nil pointer dereference
```
**Solution:** Always ensure pointers are initialized before dereferencing.

### 3. **Misusing Interfaces**
Implementing interfaces with pointer receivers when not necessary can lead to unexpected behavior.

```
type Reader interface {
    Read(p []byte) (n int, err error)
}

type File struct{}

func (f *File) Read(p []byte) (int, error) {
    // Implementation
    return 0, nil
}

func main() {
    var r Reader = File{} // Error: File does not implement Reader (Read method has pointer receiver)
}
```
**Solution:** Understand when to use pointer vs. value receivers. Implement interfaces accordingly.

### 4. **Ignoring Zero Values Leading to Unexpected Behavior**

```
type Config struct {
    Timeout time.Duration
}

func main() {
    var cfg Config
    fmt.Println(cfg.Timeout) // Outputs: 0s, which might be unintended
}
```
**Solution:** Initialize all necessary fields explicitly to prevent logic errors.
