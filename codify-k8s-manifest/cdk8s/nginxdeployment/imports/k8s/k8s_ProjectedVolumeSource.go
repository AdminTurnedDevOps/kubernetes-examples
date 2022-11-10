// k8s
package k8s


// Represents a projected volume source.
type ProjectedVolumeSource struct {
	// defaultMode are the mode bits used to set permissions on created files by default.
	//
	// Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	// sources is the list of volume projections.
	Sources *[]*VolumeProjection `field:"optional" json:"sources" yaml:"sources"`
}

