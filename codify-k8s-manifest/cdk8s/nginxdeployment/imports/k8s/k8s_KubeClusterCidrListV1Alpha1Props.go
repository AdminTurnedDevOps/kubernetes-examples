// k8s
package k8s


// ClusterCIDRList contains a list of ClusterCIDR.
type KubeClusterCidrListV1Alpha1Props struct {
	// Items is the list of ClusterCIDRs.
	Items *[]*KubeClusterCidrv1Alpha1Props `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

