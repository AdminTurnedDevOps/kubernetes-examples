// k8s
package k8s


// HorizontalPodAutoscalerList is a list of horizontal pod autoscaler objects.
type KubeHorizontalPodAutoscalerListV2Props struct {
	// items is the list of horizontal pod autoscaler objects.
	Items *[]*KubeHorizontalPodAutoscalerV2Props `field:"required" json:"items" yaml:"items"`
	// metadata is the standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

