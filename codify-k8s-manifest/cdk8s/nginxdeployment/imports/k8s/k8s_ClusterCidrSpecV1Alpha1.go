// k8s
package k8s


// ClusterCIDRSpec defines the desired state of ClusterCIDR.
type ClusterCidrSpecV1Alpha1 struct {
	// PerNodeHostBits defines the number of host bits to be configured per node.
	//
	// A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.
	PerNodeHostBits *float64 `field:"required" json:"perNodeHostBits" yaml:"perNodeHostBits"`
	// IPv4 defines an IPv4 IP block in CIDR notation(e.g. "10.0.0.0/8"). At least one of IPv4 and IPv6 must be specified. This field is immutable.
	Ipv4 *string `field:"optional" json:"ipv4" yaml:"ipv4"`
	// IPv6 defines an IPv6 IP block in CIDR notation(e.g. "fd12:3456:789a:1::/64"). At least one of IPv4 and IPv6 must be specified. This field is immutable.
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
	// NodeSelector defines which nodes the config is applicable to.
	//
	// An empty or nil NodeSelector selects all nodes. This field is immutable.
	NodeSelector *NodeSelector `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
}

