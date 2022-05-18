terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=>3.0.0"
    }
  }
}

provider "azurerm" {
  features {}
}

resource "azurerm_public_ip" "pubip" {
  name                = "pubip-lb"
  resource_group_name = var.rg_name
  location            = var.location
  allocation_method   = "Static"

}