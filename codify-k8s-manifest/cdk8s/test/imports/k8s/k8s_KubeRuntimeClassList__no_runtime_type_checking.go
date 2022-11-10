//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeRuntimeClassList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeRuntimeClassList_ManifestParameters(props *KubeRuntimeClassListProps) error {
	return nil
}

func validateKubeRuntimeClassList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeRuntimeClassListParameters(scope constructs.Construct, id *string, props *KubeRuntimeClassListProps) error {
	return nil
}

