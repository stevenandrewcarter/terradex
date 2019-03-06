terraform {
  backend "http" {
    address = "http://localhost:8080/second_source"
  }
}