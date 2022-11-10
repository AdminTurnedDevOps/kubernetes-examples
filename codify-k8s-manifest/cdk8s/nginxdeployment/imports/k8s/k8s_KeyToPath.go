// k8s
package k8s


// Maps a string key to a path within a volume.
type KeyToPath struct {
	// key is the key to project.
	Key *string `field:"required" json:"key" yaml:"key"`
	// path is the relative path of the file to map the key to.
	//
	// May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.
	Path *string `field:"required" json:"path" yaml:"path"`
	// mode is Optional: mode bits used to set permissions on this file.
	//
	// Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
}

