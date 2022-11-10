// k8s
package k8s


// IngressBackend describes all endpoints for a given service and port.
type IngressBackend struct {
	// Resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object.
	//
	// If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with "Service".
	Resource *TypedLocalObjectReference `field:"optional" json:"resource" yaml:"resource"`
	// Service references a Service as a Backend.
	//
	// This is a mutually exclusive setting with "Resource".
	Service *IngressServiceBackend `field:"optional" json:"service" yaml:"service"`
}

