// k8s
package k8s


// SelfSubjectRulesReviewSpec defines the specification for SelfSubjectRulesReview.
type SelfSubjectRulesReviewSpec struct {
	// Namespace to evaluate rules for.
	//
	// Required.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

