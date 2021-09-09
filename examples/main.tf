terraform {
  required_providers {
    toolbox = {
      version = "0.0.1"
      source  = "dihedron.org/cloud/toolbox"
    }
  }
}

provider "toolbox" {}

data "toolbox_regex" "re1" {
  pattern = ".*"
  input = "abc"
}

# module "psl" {
#   source = "./coffee"

#   coffee_name = "Packer Spiced Latte"
# }

# output "psl" {
#   value = module.psl.coffee
# }
