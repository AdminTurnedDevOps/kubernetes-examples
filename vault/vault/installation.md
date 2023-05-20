**Make sure you enable Raft so data gets replicated across all the nodes**
## PVC's are created automatically for persistent storage.

1. Ensure KMS is configured.

2. Create a secret for access to authenticate to AWS for auto-unseal
```
kubectl create secret generic -n vault eks-creds \
    --from-literal=AWS_ACCESS_KEY_ID="" \
    --from-literal=AWS_SECRET_ACCESS_KEY=""
```

3. Add the Helm repo
```
helm repo add hashicorp https://helm.releases.hashicorp.com
```

4. Run the Helm configuration
```
helm install vault hashicorp/vault \
    -f ./override-values.yaml \
    --namespace vault \
    --create-namespace
```