//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeIngressClass_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeIngressClass_ManifestParameters(props *KubeIngressClassProps) error {
	return nil
}

func validateKubeIngressClass_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeIngressClassParameters(scope constructs.Construct, id *string, props *KubeIngressClassProps) error {
	return nil
}

