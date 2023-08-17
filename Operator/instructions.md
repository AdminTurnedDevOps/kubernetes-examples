1. Create a new Directory called Operator

2. Within the Operator directory, create a new directory called `mikestestop`

3. Install Kubebuilder
```
curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
```

```
kubebuilder init --domain simplyanengineer.com --repo github.com/adminturneddevops/mikesapi
```

```
kubebuilder create api --group simplyanengineer.com --version v1 --kind MikesAPI
```

```
make manifests
```