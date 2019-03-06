terraform {
  backend "http" {
    address = "http://localhost:8080/za_kraken"
//    lock_address = "http://localhost:8080/za_kraken/lock"
//    unlock_address = "http://localhost:8080/za_kraken/unlock"
  }
}