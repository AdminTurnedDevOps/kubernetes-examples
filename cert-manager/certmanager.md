kubectl create namespace cert-manager

helm repo add jetstack https://charts.jetstack.io

helm repo update

helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --set installCRDs=true

kubectl get pods --namespace cert-manager