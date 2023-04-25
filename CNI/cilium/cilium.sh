helm repo add cilium https://helm.cilium.io/

  helm install cilium cilium/cilium \
  --namespace kube-system \
  --set egressMasqueradeInterfaces=eth0 \
  --set tunnel=disabled \
  --set eni.enabled=true \
  --set ipam.mode=eni \
  --set nodeinit.enabled=true \
  --set kubeProxyReplacement=strict

kubectl delete ds -n kube-system aws-node                       
kubectl delete ds -n kube-system kube-proxy