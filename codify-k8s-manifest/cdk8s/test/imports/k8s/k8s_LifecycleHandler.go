// k8s
package k8s


// LifecycleHandler defines a specific action that should be taken in a lifecycle hook.
//
// One and only one of the fields, except TCPSocket must be specified.
type LifecycleHandler struct {
	// Exec specifies the action to take.
	Exec *ExecAction `field:"optional" json:"exec" yaml:"exec"`
	// HTTPGet specifies the http request to perform.
	HttpGet *HttpGetAction `field:"optional" json:"httpGet" yaml:"httpGet"`
	// Deprecated.
	//
	// TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.
	TcpSocket *TcpSocketAction `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

