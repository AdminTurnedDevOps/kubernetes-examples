helm repo add hashicorp https://helm.releases.hashicorp.com

kubectl create namespace consul

helm upgrade --install -n consul consul hashicorp/consul --wait -f - <<EOF
global:
  name: consul
server:
  replicas: 1
  bootstrapExpect: 1
connectInject:
  enabled: true
EOF

curl -sL https://run.linkerd.io/emojivoto.yml \
  | sed 's|    metadata:|    metadata:\n      annotations:\n        consul.hashicorp.com/connect-inject: "true"|' \
  | sed 's|targetPort: 8080|targetPort: 20000|' \
  | kubectl apply -f -
