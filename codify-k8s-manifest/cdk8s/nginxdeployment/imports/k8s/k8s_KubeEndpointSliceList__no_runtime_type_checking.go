//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeEndpointSliceList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeEndpointSliceList_ManifestParameters(props *KubeEndpointSliceListProps) error {
	return nil
}

func validateKubeEndpointSliceList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeEndpointSliceListParameters(scope constructs.Construct, id *string, props *KubeEndpointSliceListProps) error {
	return nil
}

