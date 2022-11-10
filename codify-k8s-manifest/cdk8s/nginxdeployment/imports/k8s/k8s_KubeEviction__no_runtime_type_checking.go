//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeEviction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeEviction_ManifestParameters(props *KubeEvictionProps) error {
	return nil
}

func validateKubeEviction_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeEvictionParameters(scope constructs.Construct, id *string, props *KubeEvictionProps) error {
	return nil
}

