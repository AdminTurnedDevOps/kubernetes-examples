// k8s
package k8s


// HTTPIngressRuleValue is a list of http selectors pointing to backends.
//
// In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.
type HttpIngressRuleValue struct {
	// A collection of paths that map requests to backends.
	Paths *[]*HttpIngressPath `field:"required" json:"paths" yaml:"paths"`
}

