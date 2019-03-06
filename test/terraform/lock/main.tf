resource "null_resource" "lock" {
  triggers {
    test = "2"
  }
}

output "result_output" {
  value = "1"
}