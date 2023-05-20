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