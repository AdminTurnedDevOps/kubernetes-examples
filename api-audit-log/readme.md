`cd api-audit-log`

`mkdir ~/.minikube/files/etc`

`cp audit.yaml ~/.minikube/files/etc`

<!-- `mkdir audit && cd audit && touch audit.log` -->

`cd api-audit-log`

minikube start \
  --extra-config=apiserver.audit-policy-file=/etc/audit.yaml \
  --extra-config=apiserver.audit-log-path=-

`minikube logs -f`