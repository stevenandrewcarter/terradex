terraform {
  backend "http" {
    address = "http://localhost:8080/remote_state"
  }
}