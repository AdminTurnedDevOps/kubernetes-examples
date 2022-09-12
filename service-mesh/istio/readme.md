kubectl get nodes
curl -L https://istio.io/downloadIstio | sh
export PATH=$PWD/bin:$PATH
ls
echo "export PATH=$PATH:$HOME/istio-1.15.0/bin" >> ~/.bashrc
istioctl install --set values.gateways.istio-ingressgateway.enabled=false
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.15/samples/addons/kiali.yaml
kubectl get service -n istio-system
kubectl port-forward -n istio-system service/kiali :20001

istioctl kube-inject -f nginx.yaml | kubectl apply -f -