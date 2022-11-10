//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubePodDisruptionBudgetList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubePodDisruptionBudgetList_ManifestParameters(props *KubePodDisruptionBudgetListProps) error {
	return nil
}

func validateKubePodDisruptionBudgetList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubePodDisruptionBudgetListParameters(scope constructs.Construct, id *string, props *KubePodDisruptionBudgetListProps) error {
	return nil
}

