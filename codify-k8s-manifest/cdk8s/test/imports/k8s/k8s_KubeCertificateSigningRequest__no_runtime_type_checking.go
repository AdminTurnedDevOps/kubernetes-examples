//go:build no_runtime_type_checking

// k8s
package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeCertificateSigningRequest_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeCertificateSigningRequest_ManifestParameters(props *KubeCertificateSigningRequestProps) error {
	return nil
}

func validateKubeCertificateSigningRequest_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeCertificateSigningRequestParameters(scope constructs.Construct, id *string, props *KubeCertificateSigningRequestProps) error {
	return nil
}

