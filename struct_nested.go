package main

import "fmt"

// Address struct
type Address struct {
  Street  string
  City    string
  Country string
}

// Employee with nested Address
type Employee struct {
  Name    string
  Age     int
  Address Address
  Salary  float64
}

func main() {
  fmt.Println("=== Nested Structs ===")

  emp := Employee{
    Name: "John Smith",
    Age:  35,
    Address: Address{
      Street:  "123 Main St",
      City:    "San Francisco",
      Country: "USA",
    },
    Salary: 75000.00,
  }

  fmt.Println("Employee:", emp.Name)
  fmt.Println("Age:", emp.Age)
  fmt.Println("Street:", emp.Address.Street)
  fmt.Println("City:", emp.Address.City)
  fmt.Println("Country:", emp.Address.Country)
  fmt.Printf("Salary: $%.2f\n", emp.Salary)

  fmt.Println("\n=== Modifying Nested Fields ===")
  emp.Address.City = "Seattle"
  emp.Salary = 80000.00

  fmt.Println("New City:", emp.Address.City)
  fmt.Printf("New Salary: $%.2f\n", emp.Salary)
}
