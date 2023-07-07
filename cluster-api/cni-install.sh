```
helm repo add projectcalico https://docs.tigera.io/calico/charts --kubeconfig=./capi-azure.kubeconfig && \
helm install calico projectcalico/tigera-operator --kubeconfig=./capi-azure.kubeconfig -f https://raw.githubusercontent.com/kubernetes-sigs/cluster-api-provider-azure/main/templates/addons/calico/values.yaml --namespace tigera-operator --create-namespace
```