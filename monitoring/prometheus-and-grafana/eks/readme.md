aws eks --region region update-kubeconfig --name cluster_name

kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

kubectl get pods -n kube-system

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

helm install prometheus \
  prometheus-community/kube-prometheus-stack \
  --namespace monitoring \
  --create-namespace \
  --set alertmanager.persistentVolume.storageClass="gp2",server.persistentVolume.storageClass="gp2"

kubectl get all -n monitoring

kubectl port-forward svc/prometheus-kube-prometheus-prometheus -n monitoring 4001:9090
