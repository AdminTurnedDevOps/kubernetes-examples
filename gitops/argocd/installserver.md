`kubectl apply -f https://raw.githubusercontent.com/argoproj/argo-cd/v2.4.3/manifests/ha/install.yaml`

`kubectl get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d`

`kubectl port-forward service/argocd-server :80`

Get your current context for your k8s cluster

`kubectl config get-contexts -o name`

Add the context
`argocd cluster add kubernetes_context_name_here`

Log into the server via the CLI
`argocd login 127.0.0.1:56067`