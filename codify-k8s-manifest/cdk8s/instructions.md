
`mkdir nginxdeployment`

`cd nginxdeployment`

`cdk8s init go-app`

Change the module name in the `go.mod` file to `nginxdeployment`

Replace everything in the package to point to `nginxdeployment` instead of `example.com/nginxdeployment`

`go clean -modcache`

`go mod tidy`

`kubectl apply -f dist/nginxdeployment.k8s.yaml`