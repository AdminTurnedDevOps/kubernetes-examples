`kubectl apply -f https://raw.githubusercontent.com/argoproj/argo-cd/v2.4.3/manifests/ha/install.yaml`

`kubectl  get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d`

`kubectl port-forward service/argocd-server :80`