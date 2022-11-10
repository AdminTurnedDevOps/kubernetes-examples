ACA_PERS_RESOURCE_GROUP=techmarketing-rg
ACA_PERS_STORAGE_ACCOUNT_NAME=portainerdata
ACA_PERS_LOCATION=eastus
ACA_PERS_SHARE_NAME=portainerdata
CONTAINERAPPS_ENVIRONMENT=portainertest1

az storage account create \
    --resource-group $ACA_PERS_RESOURCE_GROUP \
    --name $ACA_PERS_STORAGE_ACCOUNT_NAME \
    --location $ACA_PERS_LOCATION \
    --sku Standard_LRS

# Create the file share
az storage share create \
  --name $ACA_PERS_SHARE_NAME \
  --account-name $ACA_PERS_STORAGE_ACCOUNT_NAME

echo $ACI_PERS_STORAGE_ACCOUNT_NAME

STORAGE_KEY=$(az storage account keys list --resource-group $ACI_PERS_RESOURCE_GROUP --account-name $ACI_PERS_STORAGE_ACCOUNT_NAME --query "[0].value" --output tsv)
echo $STORAGE_KEY

az extension add --name containerapp --upgrade

az containerapp env create \
  --name $CONTAINERAPPS_ENVIRONMENT \
  --resource-group $ACA_PERS_RESOURCE_GROUP \
  --location $ACA_PERS_LOCATION

  az containerapp env storage set --name $CONTAINERAPPS_ENVIRONMENT --resource-group $ACA_PERS_RESOURCE_GROUP \
    --storage-name portainerdata \
    --azure-file-account-name $ACA_PERS_STORAGE_ACCOUNT_NAME \
    --azure-file-account-key $STORAGE_KEY \
    --azure-file-share-name $ACA_PERS_SHARE_NAME \
    --access-mode ReadWrite

az containerapp create \
  --name portainer \
  --resource-group $ACA_PERS_RESOURCE_GROUP \
  --environment $CONTAINERAPPS_ENVIRONMENT \
  --image portainer/portainer-ce:latest \
  --target-port 9000 \
  --ingress 'external' \
  --query properties.configuration.ingress.fqdn

