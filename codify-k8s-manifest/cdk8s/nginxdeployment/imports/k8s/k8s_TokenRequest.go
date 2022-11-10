// k8s
package k8s


// TokenRequest contains parameters of a service account token.
type TokenRequest struct {
	// Audience is the intended audience of the token in "TokenRequestSpec".
	//
	// It will default to the audiences of kube apiserver.
	Audience *string `field:"required" json:"audience" yaml:"audience"`
	// ExpirationSeconds is the duration of validity of the token in "TokenRequestSpec".
	//
	// It has the same default value of "ExpirationSeconds" in "TokenRequestSpec".
	ExpirationSeconds *float64 `field:"optional" json:"expirationSeconds" yaml:"expirationSeconds"`
}

