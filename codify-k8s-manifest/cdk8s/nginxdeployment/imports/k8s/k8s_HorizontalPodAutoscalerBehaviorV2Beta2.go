// k8s
package k8s


// HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).
type HorizontalPodAutoscalerBehaviorV2Beta2 struct {
	// scaleDown is scaling policy for scaling Down.
	//
	// If not set, the default value is to allow to scale down to minReplicas pods, with a 300 second stabilization window (i.e., the highest recommendation for the last 300sec is used).
	ScaleDown *HpaScalingRulesV2Beta2 `field:"optional" json:"scaleDown" yaml:"scaleDown"`
	// scaleUp is scaling policy for scaling Up.
	//
	// If not set, the default value is the higher of:
	// * increase no more than 4 pods per 60 seconds
	// * double the number of pods per 60 seconds
	// No stabilization is used.
	ScaleUp *HpaScalingRulesV2Beta2 `field:"optional" json:"scaleUp" yaml:"scaleUp"`
}

