# Get dependent virtual network outputs to peer
data "terraform_remote_state" "dependencies" {
  backend = "local"

  config = {
    path = "../examples/dependencies/terraform.tfstate"
  }
}
