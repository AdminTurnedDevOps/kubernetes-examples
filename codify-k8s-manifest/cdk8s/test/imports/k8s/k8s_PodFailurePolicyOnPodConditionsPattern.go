// k8s
package k8s


// PodFailurePolicyOnPodConditionsPattern describes a pattern for matching an actual pod condition type.
type PodFailurePolicyOnPodConditionsPattern struct {
	// Specifies the required Pod condition status.
	//
	// To match a pod condition it is required that the specified status equals the pod condition status. Defaults to True.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Specifies the required Pod condition type.
	//
	// To match a pod condition it is required that specified type equals the pod condition type.
	Type *string `field:"required" json:"type" yaml:"type"`
}

