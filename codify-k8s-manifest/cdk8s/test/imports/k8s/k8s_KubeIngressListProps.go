// k8s
package k8s


// IngressList is a collection of Ingress.
type KubeIngressListProps struct {
	// Items is the list of Ingress.
	Items *[]*KubeIngressProps `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

