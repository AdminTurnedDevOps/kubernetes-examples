// k8s
package k8s


// NodeSpec describes the attributes that a node is created with.
type NodeSpec struct {
	// Deprecated: Previously used to specify the source of the node's configuration for the DynamicKubeletConfig feature.
	//
	// This feature is removed from Kubelets as of 1.24 and will be fully removed in 1.26.
	ConfigSource *NodeConfigSource `field:"optional" json:"configSource" yaml:"configSource"`
	// Deprecated.
	//
	// Not all kubelets will set this field. Remove field after 1.13. see: https://issues.k8s.io/61966
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// PodCIDR represents the pod IP range assigned to the node.
	PodCidr *string `field:"optional" json:"podCidr" yaml:"podCidr"`
	// podCIDRs represents the IP ranges assigned to the node for usage by Pods on that node.
	//
	// If this field is specified, the 0th entry must match the podCIDR field. It may contain at most 1 value for each of IPv4 and IPv6.
	PodCidRs *[]*string `field:"optional" json:"podCidRs" yaml:"podCidRs"`
	// ID of the node assigned by the cloud provider in the format: <ProviderName>://<ProviderSpecificNodeID>.
	ProviderId *string `field:"optional" json:"providerId" yaml:"providerId"`
	// If specified, the node's taints.
	Taints *[]*Taint `field:"optional" json:"taints" yaml:"taints"`
	// Unschedulable controls node schedulability of new pods.
	//
	// By default, node is schedulable. More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration
	Unschedulable *bool `field:"optional" json:"unschedulable" yaml:"unschedulable"`
}

