//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeCsiDriverList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeCsiDriverList_ManifestParameters(props *KubeCsiDriverListProps) error {
	return nil
}

func validateKubeCsiDriverList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeCsiDriverListParameters(scope constructs.Construct, id *string, props *KubeCsiDriverListProps) error {
	return nil
}

