// k8s
package k8s


// Secret holds secret data of a certain type.
//
// The total bytes of the values in the Data field must be less than MaxSecretSize bytes.
type KubeSecretProps struct {
	// Data contains the secret data.
	//
	// Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// Immutable, if set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// stringData allows specifying non-binary secret data in string form.
	//
	// It is provided as a write-only input field for convenience. All keys and values are merged into the data field on write, overwriting any existing values. The stringData field is never output when reading from the API.
	StringData *map[string]*string `field:"optional" json:"stringData" yaml:"stringData"`
	// Used to facilitate programmatic handling of secret data.
	//
	// More info: https://kubernetes.io/docs/concepts/configuration/secret/#secret-types
	Type *string `field:"optional" json:"type" yaml:"type"`
}

