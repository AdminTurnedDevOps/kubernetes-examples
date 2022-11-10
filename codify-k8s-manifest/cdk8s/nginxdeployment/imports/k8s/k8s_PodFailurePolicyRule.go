// k8s
package k8s


// PodFailurePolicyRule describes how a pod failure is handled when the requirements are met.
//
// One of OnExitCodes and onPodConditions, but not both, can be used in each rule.
type PodFailurePolicyRule struct {
	// Specifies the action taken on a pod failure when the requirements are satisfied.
	//
	// Possible values are: - FailJob: indicates that the pod's job is marked as Failed and all
	// running pods are terminated.
	// - Ignore: indicates that the counter towards the .backoffLimit is not
	// incremented and a replacement pod is created.
	// - Count: indicates that the pod is handled in the default way - the
	// counter towards the .backoffLimit is incremented.
	// Additional values are considered to be added in the future. Clients should react to an unknown action by skipping the rule.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Represents the requirement on the pod conditions.
	//
	// The requirement is represented as a list of pod condition patterns. The requirement is satisfied if at least one pattern matches an actual pod condition. At most 20 elements are allowed.
	OnPodConditions *[]*PodFailurePolicyOnPodConditionsPattern `field:"required" json:"onPodConditions" yaml:"onPodConditions"`
	// Represents the requirement on the container exit codes.
	OnExitCodes *PodFailurePolicyOnExitCodesRequirement `field:"optional" json:"onExitCodes" yaml:"onExitCodes"`
}

