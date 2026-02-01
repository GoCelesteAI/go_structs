package main

import "fmt"

func main() {
  fmt.Println("=== Anonymous Structs ===")

  // Anonymous struct - defined and used inline
  config := struct {
    Host    string
    Port    int
    Debug   bool
    Timeout int
  }{
    Host:    "localhost",
    Port:    8080,
    Debug:   true,
    Timeout: 30,
  }

  fmt.Println("Config:", config)
  fmt.Println("Host:", config.Host)
  fmt.Println("Port:", config.Port)
  fmt.Println("Debug:", config.Debug)

  fmt.Println("\n=== Slice of Anonymous Structs ===")
  users := []struct {
    Name  string
    Email string
  }{
    {"Alice", "alice@example.com"},
    {"Bob", "bob@example.com"},
    {"Charlie", "charlie@example.com"},
  }

  for i, user := range users {
    fmt.Printf("%d. %s <%s>\n", i+1, user.Name, user.Email)
  }

  fmt.Println("\n=== Quick Data Grouping ===")
  point := struct{ X, Y int }{10, 20}
  fmt.Printf("Point: (%d, %d)\n", point.X, point.Y)
}
