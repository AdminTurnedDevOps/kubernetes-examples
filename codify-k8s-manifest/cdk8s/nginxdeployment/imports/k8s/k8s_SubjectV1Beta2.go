// k8s
package k8s


// Subject matches the originator of a request, as identified by the request authentication system.
//
// There are three ways of matching an originator; by user, group, or service account.
type SubjectV1Beta2 struct {
	// `kind` indicates which one of the other fields is non-empty.
	//
	// Required.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// `group` matches based on user group name.
	Group *GroupSubjectV1Beta2 `field:"optional" json:"group" yaml:"group"`
	// `serviceAccount` matches ServiceAccounts.
	ServiceAccount *ServiceAccountSubjectV1Beta2 `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// `user` matches based on username.
	User *UserSubjectV1Beta2 `field:"optional" json:"user" yaml:"user"`
}

