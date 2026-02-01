# Go Lesson 15: Structs

Demo code from Go Tutorial for Beginners - Lesson 15.

## Files

- `struct_basics.go` - Defining structs, creating instances, struct literals
- `struct_modify.go` - Modifying fields, copying, comparison
- `struct_nested.go` - Nested structs, accessing nested fields
- `struct_anonymous.go` - Anonymous structs, inline definitions

## Running the Examples

```bash
# Struct basics
go run struct_basics.go

# Modifying struct fields
go run struct_modify.go

# Nested structs
go run struct_nested.go

# Anonymous structs
go run struct_anonymous.go
```

## Key Concepts

### Defining a Struct
```go
type Person struct {
    Name string
    Age  int
    City string
}
```

### Creating Instances
```go
// Named fields (preferred)
p := Person{Name: "Alice", Age: 25, City: "NYC"}

// Positional (order matters)
p := Person{"Bob", 30, "London"}

// Partial (others get zero values)
p := Person{Name: "Charlie"}
```

### Nested Structs
```go
type Address struct {
    City    string
    Country string
}

type Employee struct {
    Name    string
    Address Address
}

emp := Employee{
    Name: "John",
    Address: Address{City: "SF", Country: "USA"},
}
fmt.Println(emp.Address.City) // SF
```

### Anonymous Structs
```go
config := struct {
    Host string
    Port int
}{
    Host: "localhost",
    Port: 8080,
}
```

## Watch the Tutorial

[Go Tutorial for Beginners #15 - Structs](https://youtube.com/@CelesteAI)

## License

MIT
