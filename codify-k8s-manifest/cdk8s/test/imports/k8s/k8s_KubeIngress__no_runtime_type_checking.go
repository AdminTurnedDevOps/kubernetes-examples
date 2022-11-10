//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeIngress_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeIngress_ManifestParameters(props *KubeIngressProps) error {
	return nil
}

func validateKubeIngress_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeIngressParameters(scope constructs.Construct, id *string, props *KubeIngressProps) error {
	return nil
}

