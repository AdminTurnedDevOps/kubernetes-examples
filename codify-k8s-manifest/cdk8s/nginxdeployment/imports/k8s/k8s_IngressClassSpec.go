// k8s
package k8s


// IngressClassSpec provides information about the class of an Ingress.
type IngressClassSpec struct {
	// Controller refers to the name of the controller that should handle this class.
	//
	// This allows for different "flavors" that are controlled by the same controller. For example, you may have different Parameters for the same implementing controller. This should be specified as a domain-prefixed path no more than 250 characters in length, e.g. "acme.io/ingress-controller". This field is immutable.
	Controller *string `field:"optional" json:"controller" yaml:"controller"`
	// Parameters is a link to a custom resource containing additional configuration for the controller.
	//
	// This is optional if the controller does not require extra parameters.
	Parameters *IngressClassParametersReference `field:"optional" json:"parameters" yaml:"parameters"`
}

