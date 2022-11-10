// k8s
package k8s


// NetworkPolicyPort describes a port to allow traffic on.
type NetworkPolicyPort struct {
	// If set, indicates that the range of ports from port to endPort, inclusive, should be allowed by the policy.
	//
	// This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port.
	EndPort *float64 `field:"optional" json:"endPort" yaml:"endPort"`
	// The port on the given protocol.
	//
	// This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.
	Port IntOrString `field:"optional" json:"port" yaml:"port"`
	// The protocol (TCP, UDP, or SCTP) which traffic must match.
	//
	// If not specified, this field defaults to TCP.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

