//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeHorizontalPodAutoscalerV2_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeHorizontalPodAutoscalerV2_ManifestParameters(props *KubeHorizontalPodAutoscalerV2Props) error {
	return nil
}

func validateKubeHorizontalPodAutoscalerV2_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeHorizontalPodAutoscalerV2Parameters(scope constructs.Construct, id *string, props *KubeHorizontalPodAutoscalerV2Props) error {
	return nil
}

