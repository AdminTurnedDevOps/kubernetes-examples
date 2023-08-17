```
export CLUSTER_TOPOLOGY=true
```

```
export AZURE_SUBSCRIPTION_ID="220284d2-6a19-4781-87f8-5c564ec4fec9"

# Create an Azure Service Principal and paste the output here
export AZURE_TENANT_ID="fbbd9689-2869-42b8-92fb-db1c89b662ff"
export AZURE_CLIENT_ID="94b0aa03-e57f-46a9-a511-130271d4756d"
export AZURE_CLIENT_SECRET="WtB8Q~EUQY5dd11rb2avYC2QyUYE7SFj6M-B.cdN"

# Base64 encode the variables
export AZURE_SUBSCRIPTION_ID_B64="$(echo -n "$AZURE_SUBSCRIPTION_ID" | base64 | tr -d '\n')"
export AZURE_TENANT_ID_B64="$(echo -n "$AZURE_TENANT_ID" | base64 | tr -d '\n')"
export AZURE_CLIENT_ID_B64="$(echo -n "$AZURE_CLIENT_ID" | base64 | tr -d '\n')"
export AZURE_CLIENT_SECRET_B64="$(echo -n "$AZURE_CLIENT_SECRET" | base64 | tr -d '\n')"

# Settings needed for AzureClusterIdentity used by the AzureCluster
export AZURE_CLUSTER_IDENTITY_SECRET_NAME="cluster-identity-secret"
export CLUSTER_IDENTITY_NAME="cluster-identity"
export AZURE_CLUSTER_IDENTITY_SECRET_NAMESPACE="default"
```

```
# Create a secret to include the password of the Service Principal identity created in Azure
kubectl create secret generic "${AZURE_CLUSTER_IDENTITY_SECRET_NAME}" --from-literal=clientSecret="${AZURE_CLIENT_SECRET}" --namespace "${AZURE_CLUSTER_IDENTITY_SECRET_NAMESPACE}"
```

```
# Finally, initialize the management cluster
clusterctl init --infrastructure azure
```