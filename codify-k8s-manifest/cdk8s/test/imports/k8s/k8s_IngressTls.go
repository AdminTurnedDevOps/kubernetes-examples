// k8s
package k8s


// IngressTLS describes the transport layer security associated with an Ingress.
type IngressTls struct {
	// Hosts are a list of hosts included in the TLS certificate.
	//
	// The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// SecretName is the name of the secret used to terminate TLS traffic on port 443.
	//
	// Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

