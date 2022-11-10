// k8s
package k8s


// RuntimeClassList is a list of RuntimeClass objects.
type KubeRuntimeClassListProps struct {
	// Items is a list of schema objects.
	Items *[]*KubeRuntimeClassProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

