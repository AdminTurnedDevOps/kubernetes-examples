// k8s
package k8s


// Volume represents a named volume in a pod that may be accessed by any container in the pod.
type Volume struct {
	// name of the volume.
	//
	// Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `field:"required" json:"name" yaml:"name"`
	// awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	AwsElasticBlockStore *AwsElasticBlockStoreVolumeSource `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	// azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	AzureDisk *AzureDiskVolumeSource `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	// azureFile represents an Azure File Service mount on the host and bind mount to the pod.
	AzureFile *AzureFileVolumeSource `field:"optional" json:"azureFile" yaml:"azureFile"`
	// cephFS represents a Ceph FS mount on the host that shares a pod's lifetime.
	Cephfs *CephFsVolumeSource `field:"optional" json:"cephfs" yaml:"cephfs"`
	// cinder represents a cinder volume attached and mounted on kubelets host machine.
	//
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	Cinder *CinderVolumeSource `field:"optional" json:"cinder" yaml:"cinder"`
	// configMap represents a configMap that should populate this volume.
	ConfigMap *ConfigMapVolumeSource `field:"optional" json:"configMap" yaml:"configMap"`
	// csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).
	Csi *CsiVolumeSource `field:"optional" json:"csi" yaml:"csi"`
	// downwardAPI represents downward API about the pod that should populate this volume.
	DownwardApi *DownwardApiVolumeSource `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// emptyDir represents a temporary directory that shares a pod's lifetime.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	EmptyDir *EmptyDirVolumeSource `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// ephemeral represents a volume that is handled by a cluster storage driver.
	//
	// The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.
	//
	// Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity
	// tracking are needed,
	// c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through
	// a PersistentVolumeClaim (see EphemeralVolumeSource for more
	// information on the connection between this volume type
	// and PersistentVolumeClaim).
	//
	// Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.
	//
	// Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.
	//
	// A pod can use both types of ephemeral volumes and persistent volumes at the same time.
	Ephemeral *EphemeralVolumeSource `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	// fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	Fc *FcVolumeSource `field:"optional" json:"fc" yaml:"fc"`
	// flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
	FlexVolume *FlexVolumeSource `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	// flocker represents a Flocker volume attached to a kubelet's host machine.
	//
	// This depends on the Flocker control service being running.
	Flocker *FlockerVolumeSource `field:"optional" json:"flocker" yaml:"flocker"`
	// gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	GcePersistentDisk *GcePersistentDiskVolumeSource `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	// gitRepo represents a git repository at a particular revision.
	//
	// DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.
	GitRepo *GitRepoVolumeSource `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	// glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.
	//
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md
	Glusterfs *GlusterfsVolumeSource `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	// hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container.
	//
	// This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	HostPath *HostPathVolumeSource `field:"optional" json:"hostPath" yaml:"hostPath"`
	// iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod.
	//
	// More info: https://examples.k8s.io/volumes/iscsi/README.md
	Iscsi *IscsiVolumeSource `field:"optional" json:"iscsi" yaml:"iscsi"`
	// nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs.
	Nfs *NfsVolumeSource `field:"optional" json:"nfs" yaml:"nfs"`
	// persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	// photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine.
	PhotonPersistentDisk *PhotonPersistentDiskVolumeSource `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	// portworxVolume represents a portworx volume attached and mounted on kubelets host machine.
	PortworxVolume *PortworxVolumeSource `field:"optional" json:"portworxVolume" yaml:"portworxVolume"`
	// projected items for all in one resources secrets, configmaps, and downward API.
	Projected *ProjectedVolumeSource `field:"optional" json:"projected" yaml:"projected"`
	// quobyte represents a Quobyte mount on the host that shares a pod's lifetime.
	Quobyte *QuobyteVolumeSource `field:"optional" json:"quobyte" yaml:"quobyte"`
	// rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.
	//
	// More info: https://examples.k8s.io/volumes/rbd/README.md
	Rbd *RbdVolumeSource `field:"optional" json:"rbd" yaml:"rbd"`
	// scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	ScaleIo *ScaleIoVolumeSource `field:"optional" json:"scaleIo" yaml:"scaleIo"`
	// secret represents a secret that should populate this volume.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	Secret *SecretVolumeSource `field:"optional" json:"secret" yaml:"secret"`
	// storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.
	Storageos *StorageOsVolumeSource `field:"optional" json:"storageos" yaml:"storageos"`
	// vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine.
	VsphereVolume *VsphereVirtualDiskVolumeSource `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

