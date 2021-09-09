terraform {
  required_providers {
    toolbox = {
      version = "0.0.1"
      source  = "dihedron.org/cloud/toolbox"
    }
  }
}

provider "toolbox" {}

# module "psl" {
#   source = "./coffee"

#   coffee_name = "Packer Spiced Latte"
# }

# output "psl" {
#   value = module.psl.coffee
# }
