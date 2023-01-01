`helm repo add zerto https://zapps-helm.zerto.com/z4k/stable`

`helm repo update`

`kubectl create namespace zerto`

## Set your global configurations

```
global:
  imagePullSecret: "your_customer_pull_secret"
  authentication:
    adminUser: "zerto"
    adminPassword: "zerto"
    managementUser: "zerto"
    managementPassword: "zerto"

zkm-px:
  config:
    siteId: "your_site_id"
```

`helm install zerto zerto/z4k -f value.yaml --namespace zerto`

`kubectl get all -n zerto`

## Zerto CLI Config

`wget https://z4k.zerto.com/kubectl-zrt`

`chmod +x kubectl-zrt`

`mv kubectl-zrt zrt`

`kubectl zrt`