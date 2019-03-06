terraform {
  backend "http" {
    address = "http://localhost:8080/basic"
  }
}