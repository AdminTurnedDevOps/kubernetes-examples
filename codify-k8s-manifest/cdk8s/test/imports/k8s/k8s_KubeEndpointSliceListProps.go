// k8s
package k8s


// EndpointSliceList represents a list of endpoint slices.
type KubeEndpointSliceListProps struct {
	// List of endpoint slices.
	Items *[]*KubeEndpointSliceProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

