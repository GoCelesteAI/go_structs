package main

import "fmt"

type Product struct {
  Name  string
  Price float64
  Stock int
}

func main() {
  item := Product{
    Name:  "Laptop",
    Price: 999.99,
    Stock: 50,
  }

  fmt.Println("=== Modifying Struct Fields ===")
  fmt.Println("Before:", item)

  // Modify fields directly
  item.Price = 899.99
  item.Stock = 45

  fmt.Println("After discount:", item)

  fmt.Println("\n=== Struct Copying ===")
  itemCopy := item
  itemCopy.Price = 799.99

  fmt.Println("Original:", item.Price)
  fmt.Println("Copy:", itemCopy.Price)

  fmt.Println("\n=== Struct Comparison ===")
  p1 := Product{Name: "Phone", Price: 599.99, Stock: 100}
  p2 := Product{Name: "Phone", Price: 599.99, Stock: 100}
  p3 := Product{Name: "Phone", Price: 499.99, Stock: 100}

  fmt.Println("p1 == p2:", p1 == p2)
  fmt.Println("p1 == p3:", p1 == p3)
}
