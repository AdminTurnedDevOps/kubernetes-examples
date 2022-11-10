//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubePodDisruptionBudget_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubePodDisruptionBudget_ManifestParameters(props *KubePodDisruptionBudgetProps) error {
	return nil
}

func validateKubePodDisruptionBudget_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubePodDisruptionBudgetParameters(scope constructs.Construct, id *string, props *KubePodDisruptionBudgetProps) error {
	return nil
}

