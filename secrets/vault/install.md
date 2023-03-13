```
helm repo add hashicorp https://helm.releases.hashicorp.com
```

```
helm install vault hashicorp/vault --set='ui.enabled=true' --set='ui.serviceType=LoadBalancer' --namespace vault --create-namespace
```

```
kubectl exec --stdin=true --tty=true vault-0 -n vault -- vault operator init
```

Run the below three times with three unseal keys.
```
kubectl exec --stdin=true --tty=true vault-0 -n vault -- vault operator unseal
```