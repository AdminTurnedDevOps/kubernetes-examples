helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update

kubectl create namespace istio-system

# Install Istio
helm install istio-base istio/base -n istio-system

# Install the Istio discovery chart
helm install istiod istio/istiod -n istio-system --wait

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