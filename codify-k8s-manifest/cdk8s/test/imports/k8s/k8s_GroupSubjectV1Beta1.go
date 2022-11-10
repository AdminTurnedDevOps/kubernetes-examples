// k8s
package k8s


// GroupSubject holds detailed information for group-kind subject.
type GroupSubjectV1Beta1 struct {
	// name is the user group that matches, or "*" to match all user groups.
	//
	// See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names. Required.
	Name *string `field:"required" json:"name" yaml:"name"`
}

