
helm repo add cilium https://helm.cilium.io/

  helm install cilium cilium/cilium \
  --namespace kube-system \
  --set azure.enabled=true \
  --set eni.enabled=true \
  --set tunnel=disabled \
  --set nodeinit.enabled=true \
  --set ipam.mode=azure \
  --set masquerade=false

kubectl delete ds -n kube-system aws-node                       
kubectl delete ds -n kube-system kube-proxy