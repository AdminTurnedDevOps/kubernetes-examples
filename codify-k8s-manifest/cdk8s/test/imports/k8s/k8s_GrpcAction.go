// k8s
package k8s


type GrpcAction struct {
	// Port number of the gRPC service.
	//
	// Number must be in the range 1 to 65535.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).
	//
	// If this is not specified, the default behavior is defined by gRPC.
	Service *string `field:"optional" json:"service" yaml:"service"`
}

