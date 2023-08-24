```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws-s3
spec:
  package: xpkg.upbound.io/upbound/provider-aws-s3:v0.37.0
EOF
```

```
kubectl get providers
```

Take the below and save it to a file called `creds.txt`
```
[default]
aws_access_key_id = 
aws_secret_access_key = 
```

```
kubectl create secret \
generic aws-secret \
-n crossplane-system \
--from-file=creds=./creds.txt
```

Create a Provider Config to customize Crossplane to use the secret you just created.
```
cat <<EOF | kubectl apply -f -
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: creds
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: aws-secret
      key: creds
EOF
```

```
kubectl apply -f -
apiVersion: s3.aws.upbound.io/v1beta1
kind: Bucket
metadata:
  name: mjlbucket929211
spec:
  forProvider:
    region: us-east-1
  providerConfigRef:
    name: creds
EOF
```