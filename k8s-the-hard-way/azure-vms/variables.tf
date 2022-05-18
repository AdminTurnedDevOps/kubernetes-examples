variable "name" {
  type = string
  default = "ckavm"
}

variable "location" {
  type = string
  default = "eastus"
}

variable "rg_name" {
  type = string
  default = "ckastudy"
}

variable "subnet_id" {
  type = string
  # az network vnet subnet show -g resource_group_name -n subnet_name --vnet-name vnet_name
  default = "/subscriptions/220284d2-6a19-4781-87f8-5c564ec4fec9/resourceGroups/ckastudy/providers/Microsoft.Network/virtualNetworks/cka/subnets/default"
}
variable "password" {
  type = string
  default = "Password12!@"
  sensitive = true
}