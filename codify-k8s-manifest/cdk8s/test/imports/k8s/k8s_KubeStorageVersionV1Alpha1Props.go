// k8s
package k8s


// Storage version of a specific resource.
type KubeStorageVersionV1Alpha1Props struct {
	// Spec is an empty spec.
	//
	// It is here to comply with Kubernetes API style.
	Spec interface{} `field:"required" json:"spec" yaml:"spec"`
	// The name is <group>.<resource>.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

