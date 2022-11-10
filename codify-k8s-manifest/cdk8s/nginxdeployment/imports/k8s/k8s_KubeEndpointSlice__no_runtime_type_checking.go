//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeEndpointSlice_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeEndpointSlice_ManifestParameters(props *KubeEndpointSliceProps) error {
	return nil
}

func validateKubeEndpointSlice_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeEndpointSliceParameters(scope constructs.Construct, id *string, props *KubeEndpointSliceProps) error {
	return nil
}

