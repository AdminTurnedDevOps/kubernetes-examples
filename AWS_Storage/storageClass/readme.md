```
kubectl get storageclass
```

OR

```
kubectl get sc
```

```
resource "aws_eks_addon" "csi" {
  cluster_name = aws_eks_cluster.k8squickstart-eks.name
  addon_name   = "aws-ebs-csi-driver"
}
```