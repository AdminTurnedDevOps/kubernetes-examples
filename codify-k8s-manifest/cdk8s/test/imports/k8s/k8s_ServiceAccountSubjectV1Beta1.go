// k8s
package k8s


// ServiceAccountSubject holds detailed information for service-account-kind subject.
type ServiceAccountSubjectV1Beta1 struct {
	// `name` is the name of matching ServiceAccount objects, or "*" to match regardless of name.
	//
	// Required.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `namespace` is the namespace of matching ServiceAccount objects.
	//
	// Required.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

