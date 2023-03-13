```
az ad user create --display-name pearsonuser --password 'Password12!@' --user-principal-name pearsonuser@mlevan1992outlook.onmicrosoft.com
```

```
az role definition list \
	--query "[?contains(roleName, 'Azure Kubernetes Service RBAC')].{roleName:roleName,description:description}"
```

```
az role assignment create \
    --assignee "pearsonuser@mlevan1992outlook.onmicrosoft.com" \
    --role "Azure Kubernetes Service RBAC Reader" \
    --scope $(az aks show \
        --resource-group devrelasaservice \
        --name aksenvironment01 \
        --query id -o tsv)
```