## Specifiy Git credentials
export GITHUB_TOKEN=<your-token>
export GITHUB_USER=<your-username>

## Check Flux Installation
  `flux check --pre`

## Connect To GitHub

The command below not only connects Flux to GitHub, but it creates a new repo called `flux-fleet` where you can manage all of your clusters and repos/sources from one place.

```
flux bootstrap github \
  --owner=$GITHUB_USER \
  --repository=fleet-gitops \
  --branch=main \
  --namespace=fluxname \
  --path=./clusters/minikube \
  --personal
  ```

## Flux-fleet repo

You're going to need to perform your flux commands from the `flux-fleet` repository.

1. Clone the `flux-fleet` repository to your `localhost`
2. `cd` into the `flux-fleet` repo


## Add the repo/source to Flux
```
flux create source git nginxdeployment \
  --url=https://github.com/adminturneddevops/PearsonCourses \
  --branch=main \
  --interval=30s \
  --namespace=fluxname \
  --export > ./clusters/minikube/nginxdeployment-source.yaml
```

```
git add -A && git commit -m "Add nginxdeployment GitRepository"
git push
```

## Deploy the app
```
flux create kustomization nginxdeployment \
  --target-namespace=default \
  --source=nginxdeployment \
  --path="./kustomize" \
  --prune=true \
  --namespace=fluxname \
  --interval=5m \
  --export > ./clusters/minikube/nginxdeployment-kustomization.yaml
  ```

  ```
git add -A && git commit -m "Add nginxdeployment Kustomization"
git push
```

## Watch the release
`flux get kustomizations -n fluxname --watch`