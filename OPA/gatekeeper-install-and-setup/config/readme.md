The config that you see in `config.yaml` is your definition of what Gatekeeper is allowed to create policies for. For example, in the `config.yaml`, Gatekeeper knows that it can only specifiy Policies for Pods, but no other Kubernetes resources.

The config also ensures that the policy is set across clusters, not just the node where Gatekeeper was installed.

For example, if you have Pods running on multiple worker nodes, you'd need a Config so Gatekeeper can have access to the Pods across worker nodes.