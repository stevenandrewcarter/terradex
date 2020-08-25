resource "null_resource" "basic" {
  triggers = {
    test = "2"
  }
}

output "result_output" {
  value = "1"
}