//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeHorizontalPodAutoscalerListV2_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeHorizontalPodAutoscalerListV2_ManifestParameters(props *KubeHorizontalPodAutoscalerListV2Props) error {
	return nil
}

func validateKubeHorizontalPodAutoscalerListV2_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeHorizontalPodAutoscalerListV2Parameters(scope constructs.Construct, id *string, props *KubeHorizontalPodAutoscalerListV2Props) error {
	return nil
}

