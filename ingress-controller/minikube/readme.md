`minikube start`

`minikube addons enable ingress`

`kubectl get pods -n ingress-nginx`

`kubectl apply -f deployment.yaml`

`kubectl create -f ingress.yaml`

`minikube tunnel`

OR

`kubectl port-forward service/ingress-nginx-controller -n ingress-nginx :8080`

