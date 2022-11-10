// k8s
package k8s


// ForZone provides information about which zones should consume this endpoint.
type ForZone struct {
	// name represents the name of the zone.
	Name *string `field:"required" json:"name" yaml:"name"`
}

