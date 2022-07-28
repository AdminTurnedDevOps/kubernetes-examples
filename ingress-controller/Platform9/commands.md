helm repo add nginx-stable https://helm.nginx.com/stable
helm repo update

helm install ingress-nginx nginx-stable/nginx-ingress

kubectl get pods --all-namespaces -l app.kubernetes.io/name=ingress-nginx

kubectl get services ingress-nginx-controller --namespace=ingress-nginx