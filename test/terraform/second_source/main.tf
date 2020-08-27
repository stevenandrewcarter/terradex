resource "null_resource" "second_source" {
  triggers = {
    test = "2"
  }
}
