// k8s
package k8s


// EventList is a list of Event objects.
type KubeEventListProps struct {
	// items is a list of schema objects.
	Items *[]*KubeEventProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

