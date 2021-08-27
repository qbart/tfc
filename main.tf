terraform {
  backend "http" {
    address = "http://localhost:8080/state"
    lock_address = "http://localhost:8080/lock"
    unlock_address = "http://localhost:8080/unlock"
  }
}
