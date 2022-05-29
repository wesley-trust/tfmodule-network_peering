data "terraform_remote_state" "dependencies" {
  backend = "local"

  config = {
    path = "./dependencies/terraform.tfstate"
  }
}
