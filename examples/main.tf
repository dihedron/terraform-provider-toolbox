terraform {
  required_providers {
    toolbox = {
      version = "0.0.1"
      source  = "dihedron.org/cloud/toolbox"
    }
  }
}

provider "toolbox" {}

data "toolbox_regex" "matched_regex" {
  pattern = ".*"
  input = "abc"
}

output "matched_regex_result" {
   value = data.toolbox_regex.matched_regex.matched
}

data "toolbox_regex" "unmatched_regex" {
  pattern = "^b.*"
  input = "abc"
}

output "unmatched_regex_result" {
   value = data.toolbox_regex.unmatched_regex.matched
}



# module "psl" {
#   source = "./coffee"

#   coffee_name = "Packer Spiced Latte"
# }

# output "psl" {
#   value = module.psl.coffee
# }
