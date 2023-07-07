```
helm install --kubeconfig=./capi-azure.kubeconfig --repo https://raw.githubusercontent.com/kubernetes-sigs/cloud-provider-azure/master/helm/repo cloud-provider-azure --generate-name --set infra.clusterName=capi-azure --set cloudControllerManager.clusterCIDR="192.168.0.0/16"
```