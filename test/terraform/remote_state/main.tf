data "terraform_remote_state" "basic" {
  backend = "http"
  config {
    address = "http://localhost:8080/basic"
  }
}

resource "null_resource" "remote_state_changed" {
  triggers {
    remote_state = "${data.terraform_remote_state.basic.result_output}"
  }
}