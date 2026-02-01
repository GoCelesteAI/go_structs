package main

import "fmt"

// Define a struct type
type Person struct {
  Name string
  Age  int
  City string
}

func main() {
  // Create a struct instance
  var p1 Person
  p1.Name = "Alice"
  p1.Age = 25
  p1.City = "New York"

  fmt.Println("=== Basic Struct ===")
  fmt.Println("Person:", p1)
  fmt.Println("Name:", p1.Name)
  fmt.Println("Age:", p1.Age)
  fmt.Println("City:", p1.City)

  // Struct literal - named fields
  p2 := Person{
    Name: "Bob",
    Age:  30,
    City: "London",
  }
  fmt.Println("\nNamed fields:", p2)

  // Struct literal - positional (order matters!)
  p3 := Person{"Charlie", 35, "Tokyo"}
  fmt.Println("Positional:", p3)

  // Partial initialization (missing fields get zero values)
  p4 := Person{Name: "Diana"}
  fmt.Println("Partial:", p4)
}
