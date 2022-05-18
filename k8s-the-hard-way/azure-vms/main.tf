terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.0.0"
    }
  }
}

provider "azurerm" {
  features {}
}

resource "azurerm_public_ip" "pubip" {
  count               = 4
  name                = "pubip-${count.index}"
  resource_group_name = var.rg_name
  location            = var.location
  allocation_method   = "Static"

}

resource "azurerm_network_interface" "main" {
  count               = 4
  name                = "ckanic-${count.index}"
  location            = var.location
  resource_group_name = var.rg_name

  ip_configuration {
    name                          = "testconfiguration1"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.pubip[count.index].id
  }
}

resource "azurerm_virtual_machine" "main" {
  count                 = 4
  name                  = "${var.name}-${count.index}"
  location              = var.location
  resource_group_name   = var.rg_name
  network_interface_ids = [azurerm_network_interface.main[count.index].id]
#   network_interface_ids = [element(azurerm_network_interface.main.*.id, count.index)]
  vm_size               = "Standard_B1ms"

  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "18.04-LTS"
    version   = "latest"
  }
  storage_os_disk {
    name              = "storage-${count.index}"
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }
  os_profile {
    computer_name  = "ckastudy-${count.index}"
    admin_username = "testadmin"
    admin_password = var.password
  }
  os_profile_linux_config {
    disable_password_authentication = false
  }

}
