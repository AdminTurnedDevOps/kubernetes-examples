```
export AZURE_LOCATION="eastus"

export AZURE_CONTROL_PLANE_MACHINE_TYPE="Standard_D2s_v3"
export AZURE_NODE_MACHINE_TYPE="Standard_D2s_v3"

export AZURE_RESOURCE_GROUP="devrelasaservice"
```

```
clusterctl generate cluster capi-azure --kubernetes-version v1.27.0 > capi-azurekubeadm.yaml
```