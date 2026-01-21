terraform {
    required_version = ">= 1.0.0"
}

provider "local" {}

resource "local_file" "hello" {
    content = "Hello Souf - Terraform is working"
    filename = "${path.module}/hello.txt"
}