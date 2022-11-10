//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeCsiStorageCapacityList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeCsiStorageCapacityList_ManifestParameters(props *KubeCsiStorageCapacityListProps) error {
	return nil
}

func validateKubeCsiStorageCapacityList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeCsiStorageCapacityListParameters(scope constructs.Construct, id *string, props *KubeCsiStorageCapacityListProps) error {
	return nil
}

