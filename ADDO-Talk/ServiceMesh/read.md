## Install Prometheus

You can install the Prometheus addon for Istio with the below.

```
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.19/samples/addons/prometheus.yaml
```

## Install Istio
helm repo add istio https://istio-release.storage.googleapis.com/charts

helm repo update

helm install istio-base istio/base -n istio-system --create-namespace

# Install the Istio discovery chart
helm install istiod istio/istiod -n istio-system

# Ingress
kubectl create namespace istio-ingress

kubectl label namespace istio-ingress istio-injection=enabled

helm install istio-ingress istio/gateway -n istio-ingress --wait

# Confirm status
helm status istiod -n istio-system

# Install Kiali

```
helm repo add kiali https://kiali.org/helm-charts
helm repo update

```

```
helm install \
    --set cr.create=true \
    --set cr.namespace=istio-system \
    --namespace kiali-operator \
    --create-namespace \
    kiali-operator \
    kiali/kiali-operator
```

```
kubectl port-forward svc/kiali -n istio-system 8080:20001
```

```
kubectl -n istio-system create token kiali-service-account
```

## Demo App
git clone https://github.com/microservices-demo/microservices-demo
cd microservices-demo
cd deploy/kubernetes

kubectl create namespace sock-shop

kubectl label namespace sock-shop istio-injection=enabled

kubectl apply -f complete-demo.yaml

