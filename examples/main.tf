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
  pattern = "a(x*)b"
  input = "-axxb-ab-"
}

output "regex_re1_matched" {
   value = data.toolbox_regex.re1.matched
}

output "regex_re1_matches" {
   value = data.toolbox_regex.re1.matches
}

/*
data "toolbox_regex" "unmatched_regex" {
  pattern = "^b.*"
  input = "abc"
}

output "unmatched_regex_result" {
   value = data.toolbox_regex.unmatched_regex.matched
}
*/

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

data "toolbox_cidr" "set2" {
  added = [
    "10.0.0.128/25",
    "10.0.0.32/27",
    "10.0.0.64/26",
    "10.0.1.0/24",
    "10.0.128.0/17",
    "10.0.16.0/20",
    "10.0.2.0/23",
    "10.0.32.0/19",
    "10.0.4.0/22",
    "10.0.64.0/18",
    "10.0.8.0/21",
    "10.1.0.0/16",
    "10.128.0.0/9",
    "10.16.0.0/12",
    "10.2.0.0/15",
    "10.32.0.0/11",
    "10.4.0.0/14",
    "10.64.0.0/10",
    "10.8.0.0/13",
    "172.16.0.0/12",
    "192.168.0.0/16",    
    "10.0.0.0/27"
  ]
}

output "cidr_set2" {
   value = data.toolbox_cidr.set2.prefixes
}
