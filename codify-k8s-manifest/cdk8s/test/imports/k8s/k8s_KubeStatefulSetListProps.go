// k8s
package k8s


// StatefulSetList is a collection of StatefulSets.
type KubeStatefulSetListProps struct {
	// Items is the list of stateful sets.
	Items *[]*KubeStatefulSetProps `field:"required" json:"items" yaml:"items"`
	// Standard list's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

