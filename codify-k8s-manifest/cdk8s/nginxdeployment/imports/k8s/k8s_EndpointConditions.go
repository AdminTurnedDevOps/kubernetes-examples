// k8s
package k8s


// EndpointConditions represents the current condition of an endpoint.
type EndpointConditions struct {
	// ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint.
	//
	// A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready. For compatibility reasons, ready should never be "true" for terminating endpoints.
	Ready *bool `field:"optional" json:"ready" yaml:"ready"`
	// serving is identical to ready except that it is set regardless of the terminating state of endpoints.
	//
	// This condition should be set to true for a ready endpoint that is terminating. If nil, consumers should defer to the ready condition. This field can be enabled with the EndpointSliceTerminatingCondition feature gate.
	Serving *bool `field:"optional" json:"serving" yaml:"serving"`
	// terminating indicates that this endpoint is terminating.
	//
	// A nil value indicates an unknown state. Consumers should interpret this unknown state to mean that the endpoint is not terminating. This field can be enabled with the EndpointSliceTerminatingCondition feature gate.
	Terminating *bool `field:"optional" json:"terminating" yaml:"terminating"`
}

