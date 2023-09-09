## Install Kube-Prometheus

```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```

```
helm repo update
```

```
helm install kube-prometheus -n monitoring prometheus-community/kube-prometheus-stack --create-namespace
```

```
kubectl create -f manifests/setup

until kubectl get servicemonitors --all-namespaces ; do date; sleep 1; echo ""; done

kubectl create -f manifests/
```

Grafana Creds
- Username: admin
- Password: admin

## Install Loki

```
helm repo add grafana https://grafana.github.io/helm-charts
```

Configure a `values.yaml` file
```
loki:
  commonConfig:
    replication_factor: 1
  storage:
    type: 'filesystem'
singleBinary:
  replicas: 1
```

```
helm install -n monitoring --values loki-values.yaml loki grafana/loki
```
## Install Tempo

```
helm repo add grafana https://grafana.github.io/helm-charts
```

```
helm install -n monitoring tempo grafana/tempo
```