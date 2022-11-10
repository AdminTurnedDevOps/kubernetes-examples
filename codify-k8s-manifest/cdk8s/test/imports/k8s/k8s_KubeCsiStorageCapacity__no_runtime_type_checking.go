//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeCsiStorageCapacity_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeCsiStorageCapacity_ManifestParameters(props *KubeCsiStorageCapacityProps) error {
	return nil
}

func validateKubeCsiStorageCapacity_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeCsiStorageCapacityParameters(scope constructs.Construct, id *string, props *KubeCsiStorageCapacityProps) error {
	return nil
}

