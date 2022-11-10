// k8s
package k8s


// CSIDriver captures information about a Container Storage Interface (CSI) volume driver deployed on the cluster.
//
// Kubernetes attach detach controller uses this object to determine whether attach is required. Kubelet uses this object to determine whether pod information needs to be passed on mount. CSIDriver objects are non-namespaced.
type KubeCsiDriverProps struct {
	// Specification of the CSI Driver.
	Spec *CsiDriverSpec `field:"required" json:"spec" yaml:"spec"`
	// Standard object metadata.
	//
	// metadata.Name indicates the name of the CSI driver that this object refers to; it MUST be the same name returned by the CSI GetPluginName() call for that driver. The driver name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

