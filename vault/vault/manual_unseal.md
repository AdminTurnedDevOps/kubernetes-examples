```
kubectl exec --stdin=true --tty=true vault-0 -n vault -- vault operator init
```

Run the below three times with three unseal keys.
```
kubectl exec --stdin=true --tty=true vault-0 -n vault -- vault operator unseal
```

kubectl exec -ti vault-1 -n vault -- vault operator raft join http://vault-0.vault-internal:8200
kubectl exec -ti vault-1 -n vault -- vault operator unseal

kubectl exec -ti vault-2 -n vault -- vault operator raft join http://vault-0.vault-internal:8200
kubectl exec -ti vault-2 -n vault -- vault operator unseal