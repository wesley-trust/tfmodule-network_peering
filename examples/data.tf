data "terraform_remote_state" "dependencies" {
  backend = "local"

  config = {
    path = "../examples/dependencies/terraform.tfstate"
  }
}
