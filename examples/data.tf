data "terraform_remote_state" "outputs" {
  backend = "local"

  config = {
    path = "./dependencies/terraform.tfstate"
  }
}
