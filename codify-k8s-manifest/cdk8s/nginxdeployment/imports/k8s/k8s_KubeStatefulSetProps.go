// k8s
package k8s


// StatefulSet represents a set of pods with consistent identities.
//
// Identities are defined as:
// - Network: A single stable DNS and hostname.
// - Storage: As many VolumeClaims as requested.
//
// The StatefulSet guarantees that a given network identity will always map to the same storage identity.
type KubeStatefulSetProps struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Spec defines the desired identities of pods in this set.
	Spec *StatefulSetSpec `field:"optional" json:"spec" yaml:"spec"`
}

