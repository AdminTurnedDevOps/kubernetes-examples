//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeIngressClassList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeIngressClassList_ManifestParameters(props *KubeIngressClassListProps) error {
	return nil
}

func validateKubeIngressClassList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeIngressClassListParameters(scope constructs.Construct, id *string, props *KubeIngressClassListProps) error {
	return nil
}

