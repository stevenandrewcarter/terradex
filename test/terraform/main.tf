data "terraform_remote_state" "foo" {
  backend = "http"
  config {
    address = "http://localhost:8080"
  }
}

resource "null_resource" "test" {

}