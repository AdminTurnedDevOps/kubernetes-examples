// k8s
package k8s


// CustomResourceDefinitionVersion describes a version for CRD.
type CustomResourceDefinitionVersion struct {
	// name is the version name, e.g. “v1”, “v2beta1”, etc. The custom resources are served under this version at `/apis/<group>/<version>/...` if `served` is true.
	Name *string `field:"required" json:"name" yaml:"name"`
	// served is a flag enabling/disabling this version from being served via REST APIs.
	Served *bool `field:"required" json:"served" yaml:"served"`
	// storage indicates this version should be used when persisting custom resources to storage.
	//
	// There must be exactly one version with storage=true.
	Storage *bool `field:"required" json:"storage" yaml:"storage"`
	// additionalPrinterColumns specifies additional columns returned in Table output.
	//
	// See https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables for details. If no columns are specified, a single column displaying the age of the custom resource is used.
	AdditionalPrinterColumns *[]*CustomResourceColumnDefinition `field:"optional" json:"additionalPrinterColumns" yaml:"additionalPrinterColumns"`
	// deprecated indicates this version of the custom resource API is deprecated.
	//
	// When set to true, API requests to this version receive a warning header in the server response. Defaults to false.
	Deprecated *bool `field:"optional" json:"deprecated" yaml:"deprecated"`
	// deprecationWarning overrides the default warning returned to API clients.
	//
	// May only be set when `deprecated` is true. The default warning indicates this version is deprecated and recommends use of the newest served version of equal or greater stability, if one exists.
	DeprecationWarning *string `field:"optional" json:"deprecationWarning" yaml:"deprecationWarning"`
	// schema describes the schema used for validation, pruning, and defaulting of this version of the custom resource.
	Schema *CustomResourceValidation `field:"optional" json:"schema" yaml:"schema"`
	// subresources specify what subresources this version of the defined custom resource have.
	Subresources *CustomResourceSubresources `field:"optional" json:"subresources" yaml:"subresources"`
}

