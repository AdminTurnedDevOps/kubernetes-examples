// k8s
package k8s


// APIServiceList is a list of APIService objects.
type KubeApiServiceListProps struct {
	// Items is the list of APIService.
	Items *[]*KubeApiServiceProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

