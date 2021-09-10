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

data "toolbox_cidr" "set1" {
  added = [
      "10.0.0.0/8",
      "172.16.0.0/12",
      "192.168.0.0/16"
  ]

  subtracted = [ 
    "10.0.0.0/27"
  ]
}

output "cidr_set1" {
   value = data.toolbox_cidr.set1.prefixes
}


# module "psl" {
#   source = "./coffee"

#   coffee_name = "Packer Spiced Latte"
# }

# output "psl" {
#   value = module.psl.coffee
# }
