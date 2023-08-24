```
helm repo add crossplane-stable https://charts.crossplane.io/stable
```

```
helm repo update
```

```
helm install crossplane \
crossplane-stable/crossplane \
--namespace crossplane-system \
--create-namespace
```

