# 10. Best Practices for Variables and Types

### 1. **Prefer Short Declarations for Local Variables**
Use `:=` for brevity and readability within functions.

```
count := 10
```

### 2. **Use Explicit Types for Package-Level Variables**
Enhances clarity and prevents unintended type inferences.

```
var APIEndpoint string = "https://api.example.com"
```

### 3. **Leverage Custom Types for Clarity and Safety**
Distinguish between similar data types to prevent logical errors.

```
type UserID int
type OrderID int
```

### 4. **Minimize Use of Global Variables**
Restrict scope to enhance modularity and reduce dependencies.

### 5. **Use Constants for Immutable Values**
Ensure values that shouldn't change remain constant.

```
const Pi = 3.14159
```

### 6. **Handle Zero Values Appropriately**
Initialize variables explicitly when zero values are insufficient.

### 7. **Avoid Shadowing Variables**
Prevent confusion by using distinct names across scopes.

### 8. **Use Interfaces to Define Behavior**
Promote decoupling and flexibility in your codebase.

### 9. **Optimize Memory Usage with Pointers**
Use pointers for large structs and when mutation is required.

### 10. **Document Custom Types and Methods**
Enhance maintainability and understanding of your code.

```
// Celsius represents temperature in degrees Celsius.
type Celsius float64
```




