storageAccountName='test92'
resourceGroup='devrelasaservice'

storageKey=$(az storage account keys list --resource-group $resourceGroup --account-name $storageAccountName --query "[0].value" -o tsv)
echo $storageKey

# Azure Secret
kubectl create secret generic azure-secret --from-literal=azurestorageaccountname=$storageAccountName --from-literal=azurestorageaccountkey=$storageKey