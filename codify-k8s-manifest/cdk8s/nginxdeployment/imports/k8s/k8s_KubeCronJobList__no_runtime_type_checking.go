//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeCronJobList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeCronJobList_ManifestParameters(props *KubeCronJobListProps) error {
	return nil
}

func validateKubeCronJobList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeCronJobListParameters(scope constructs.Construct, id *string, props *KubeCronJobListProps) error {
	return nil
}

