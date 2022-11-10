// k8s
package k8s


// CertificateSigningRequestList is a collection of CertificateSigningRequest objects.
type KubeCertificateSigningRequestListProps struct {
	// items is a collection of CertificateSigningRequest objects.
	Items *[]*KubeCertificateSigningRequestProps `field:"required" json:"items" yaml:"items"`
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

