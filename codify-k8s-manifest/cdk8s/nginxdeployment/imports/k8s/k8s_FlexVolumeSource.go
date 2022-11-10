// k8s
package k8s


// FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
type FlexVolumeSource struct {
	// driver is the name of the driver to use for this volume.
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// fsType is the filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// options is Optional: this field holds extra command options if any.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// readOnly is Optional: defaults to false (read/write).
	//
	// ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts.
	//
	// This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.
	SecretRef *LocalObjectReference `field:"optional" json:"secretRef" yaml:"secretRef"`
}

