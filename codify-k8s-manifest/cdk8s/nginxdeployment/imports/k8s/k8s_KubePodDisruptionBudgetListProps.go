// k8s
package k8s


// PodDisruptionBudgetList is a collection of PodDisruptionBudgets.
type KubePodDisruptionBudgetListProps struct {
	// Items is a list of PodDisruptionBudgets.
	Items *[]*KubePodDisruptionBudgetProps `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

