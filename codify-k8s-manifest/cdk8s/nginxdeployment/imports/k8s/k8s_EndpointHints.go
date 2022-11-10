// k8s
package k8s


// EndpointHints provides hints describing how an endpoint should be consumed.
type EndpointHints struct {
	// forZones indicates the zone(s) this endpoint should be consumed by to enable topology aware routing.
	ForZones *[]*ForZone `field:"optional" json:"forZones" yaml:"forZones"`
}

