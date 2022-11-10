// k8s
package k8s


// ConfigMapNodeConfigSource contains the information to reference a ConfigMap as a config source for the Node.
//
// This API is deprecated since 1.22: https://git.k8s.io/enhancements/keps/sig-node/281-dynamic-kubelet-configuration
type ConfigMapNodeConfigSource struct {
	// KubeletConfigKey declares which key of the referenced ConfigMap corresponds to the KubeletConfiguration structure This field is required in all cases.
	KubeletConfigKey *string `field:"required" json:"kubeletConfigKey" yaml:"kubeletConfigKey"`
	// Name is the metadata.name of the referenced ConfigMap. This field is required in all cases.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace is the metadata.namespace of the referenced ConfigMap. This field is required in all cases.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// ResourceVersion is the metadata.ResourceVersion of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status.
	ResourceVersion *string `field:"optional" json:"resourceVersion" yaml:"resourceVersion"`
	// UID is the metadata.UID of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status.
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
}

