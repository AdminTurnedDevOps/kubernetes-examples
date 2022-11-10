// k8s
package k8s


// Represents storage that is managed by an external CSI volume driver (Beta feature).
type CsiPersistentVolumeSource struct {
	// driver is the name of the driver to use for this volume.
	//
	// Required.
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// volumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls.
	//
	// Required.
	VolumeHandle *string `field:"required" json:"volumeHandle" yaml:"volumeHandle"`
	// controllerExpandSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerExpandVolume call.
	//
	// This is an beta field and requires enabling ExpandCSIVolumes feature gate. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
	ControllerExpandSecretRef *SecretReference `field:"optional" json:"controllerExpandSecretRef" yaml:"controllerExpandSecretRef"`
	// controllerPublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerPublishVolume and ControllerUnpublishVolume calls.
	//
	// This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
	ControllerPublishSecretRef *SecretReference `field:"optional" json:"controllerPublishSecretRef" yaml:"controllerPublishSecretRef"`
	// fsType to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// nodeExpandSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodeExpandVolume call.
	//
	// This is an alpha field and requires enabling CSINodeExpandSecret feature gate. This field is optional, may be omitted if no secret is required. If the secret object contains more than one secret, all secrets are passed.
	NodeExpandSecretRef *SecretReference `field:"optional" json:"nodeExpandSecretRef" yaml:"nodeExpandSecretRef"`
	// nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls.
	//
	// This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
	NodePublishSecretRef *SecretReference `field:"optional" json:"nodePublishSecretRef" yaml:"nodePublishSecretRef"`
	// nodeStageSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodeStageVolume and NodeStageVolume and NodeUnstageVolume calls.
	//
	// This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
	NodeStageSecretRef *SecretReference `field:"optional" json:"nodeStageSecretRef" yaml:"nodeStageSecretRef"`
	// readOnly value to pass to ControllerPublishVolumeRequest.
	//
	// Defaults to false (read/write).
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// volumeAttributes of the volume to publish.
	VolumeAttributes *map[string]*string `field:"optional" json:"volumeAttributes" yaml:"volumeAttributes"`
}

