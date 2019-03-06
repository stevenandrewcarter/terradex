terraform {
  backend "http" {
    address = "http://localhost:8080/auth"
    username = "username"
    password = "password"
  }
}