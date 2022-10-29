az aks get-credentials -n name_of_k8s_cluster -g resource_group_name 

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

helm repo update

helm install prometheus \
  prometheus-community/kube-prometheus-stack \
  --namespace monitoring \
  --create-namespace

kubectl get all -n monitoring

kubectl port-forward svc/prometheus-grafana -n monitoring 4000:80
kubectl port-forward svc/prometheus-kube-prometheus-prometheus -n monitoring 4001:9090