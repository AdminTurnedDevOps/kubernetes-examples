// k8s
package k8s


// Projection that may be projected along with other supported volume types.
type VolumeProjection struct {
	// configMap information about the configMap data to project.
	ConfigMap *ConfigMapProjection `field:"optional" json:"configMap" yaml:"configMap"`
	// downwardAPI information about the downwardAPI data to project.
	DownwardApi *DownwardApiProjection `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// secret information about the secret data to project.
	Secret *SecretProjection `field:"optional" json:"secret" yaml:"secret"`
	// serviceAccountToken is information about the serviceAccountToken data to project.
	ServiceAccountToken *ServiceAccountTokenProjection `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

