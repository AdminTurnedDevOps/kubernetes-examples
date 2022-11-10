//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeRuntimeClass_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeRuntimeClass_ManifestParameters(props *KubeRuntimeClassProps) error {
	return nil
}

func validateKubeRuntimeClass_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeRuntimeClassParameters(scope constructs.Construct, id *string, props *KubeRuntimeClassProps) error {
	return nil
}

