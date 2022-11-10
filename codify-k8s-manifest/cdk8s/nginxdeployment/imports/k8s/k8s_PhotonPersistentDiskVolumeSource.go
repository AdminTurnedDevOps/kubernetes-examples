// k8s
package k8s


// Represents a Photon Controller persistent disk resource.
type PhotonPersistentDiskVolumeSource struct {
	// pdID is the ID that identifies Photon Controller persistent disk.
	PdId *string `field:"required" json:"pdId" yaml:"pdId"`
	// fsType is the filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
}

