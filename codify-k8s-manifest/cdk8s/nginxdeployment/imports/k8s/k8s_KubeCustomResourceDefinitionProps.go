// k8s
package k8s


// CustomResourceDefinition represents a resource that should be exposed on the API server.
//
// Its name MUST be in the format <.spec.name>.<.spec.group>.
type KubeCustomResourceDefinitionProps struct {
	// spec describes how the user wants the resources to appear.
	Spec *CustomResourceDefinitionSpec `field:"required" json:"spec" yaml:"spec"`
	// Standard object's metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

