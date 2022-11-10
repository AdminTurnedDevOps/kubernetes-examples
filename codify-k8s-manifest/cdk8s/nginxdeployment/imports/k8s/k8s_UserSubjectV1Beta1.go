// k8s
package k8s


// UserSubject holds detailed information for user-kind subject.
type UserSubjectV1Beta1 struct {
	// `name` is the username that matches, or "*" to match all usernames.
	//
	// Required.
	Name *string `field:"required" json:"name" yaml:"name"`
}

