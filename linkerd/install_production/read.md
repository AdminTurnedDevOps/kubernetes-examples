
```
helm fetch --untar linkerd/linkerd-control-plane
```

```
helm install linkerd-control-plane \
  --set-file identityTrustAnchorsPEM=ca.crt \
  --set-file identity.issuer.tls.crtPEM=issuer.crt \
  --set-file identity.issuer.tls.keyPEM=issuer.key \
  -f linkerd-control-plane/values-ha.yaml \
  linkerd/linkerd-control-plane
```