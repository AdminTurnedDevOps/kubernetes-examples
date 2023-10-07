```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: upbound-provider-azure
spec:
  package: xpkg.upbound.io/upbound/provider-azure:v0.29.0
EOF
```

```
kubectl get providers
```

Run the following, copy the output, and save it to a file called `azure.json`
```
az ad sp create-for-rbac \
--sdk-auth \
--role Owner \
--scopes /subscriptions/your_sub_id
```

```
kubectl create secret \
generic azure-secret \
-n crossplane-system \
--from-file=creds=./azure.json
```

Create the Provider Config below that updates Crossplane to use the new secret that you just created for authentication/authorization to Azure.
```
cat <<EOF | kubectl apply -f -
apiVersion: azure.upbound.io/v1beta1
metadata:
  name: default
kind: ProviderConfig
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: azure-secret
      key: creds
EOF
```

```
cat <<EOF | kubectl create -f -
apiVersion: network.azure.upbound.io/v1beta1
kind: VirtualNetwork
metadata:
  name: vnet
spec:
  forProvider:
    addressSpace:
      - 10.0.0.0/16
    location: "US East"
    resourceGroupName: devrel-as-a-service
EOF
```