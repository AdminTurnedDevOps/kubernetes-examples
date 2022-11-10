// k8s
package k8s


// IngressClassParametersReference identifies an API object.
//
// This can be used to specify a cluster or namespace-scoped resource.
type IngressClassParametersReference struct {
	// Kind is the type of resource being referenced.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name is the name of resource being referenced.
	Name *string `field:"required" json:"name" yaml:"name"`
	// APIGroup is the group for the resource being referenced.
	//
	// If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
	ApiGroup *string `field:"optional" json:"apiGroup" yaml:"apiGroup"`
	// Namespace is the namespace of the resource being referenced.
	//
	// This field is required when scope is set to "Namespace" and must be unset when scope is set to "Cluster".
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Scope represents if this refers to a cluster or namespace scoped resource.
	//
	// This may be set to "Cluster" (default) or "Namespace".
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

