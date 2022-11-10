// k8s
package k8s


// TokenRequest requests a token for a given service account.
type KubeTokenRequestProps struct {
	// Spec holds information about the request being evaluated.
	Spec *TokenRequestSpec `field:"required" json:"spec" yaml:"spec"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

