terraform {
  backend "http" {
    address = "http://localhost:8080/lock"
    lock_address = "http://localhost:8080/lock"
    unlock_address = "http://localhost:8080/lock"
    username = "username"
    password = "password"
  }
}