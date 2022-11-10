// k8s
package k8s


// ValidationRule describes a validation rule written in the CEL expression language.
type ValidationRule struct {
	// Rule represents the expression which will be evaluated by CEL.
	//
	// ref: https://github.com/google/cel-spec The Rule is scoped to the location of the x-kubernetes-validations extension in the schema. The `self` variable in the CEL expression is bound to the scoped value. Example: - Rule scoped to the root of a resource with a status subresource: {"rule": "self.status.actual <= self.spec.maxDesired"}
	//
	// If the Rule is scoped to an object with properties, the accessible properties of the object are field selectable via `self.field` and field presence can be checked via `has(self.field)`. Null valued fields are treated as absent fields in CEL expressions. If the Rule is scoped to an object with additionalProperties (i.e. a map) the value of the map are accessible via `self[mapKey]`, map containment can be checked via `mapKey in self` and all entries of the map are accessible via CEL macros and functions such as `self.all(...)`. If the Rule is scoped to an array, the elements of the array are accessible via `self[i]` and also by macros and functions. If the Rule is scoped to a scalar, `self` is bound to the scalar value. Examples: - Rule scoped to a map of objects: {"rule": "self.components['Widget'].priority < 10"} - Rule scoped to a list of integers: {"rule": "self.values.all(value, value >= 0 && value < 100)"} - Rule scoped to a string value: {"rule": "self.startsWith('kube')"}
	//
	// The `apiVersion`, `kind`, `metadata.name` and `metadata.generateName` are always accessible from the root of the object and from any x-kubernetes-embedded-resource annotated objects. No other metadata properties are accessible.
	//
	// Unknown data preserved in custom resources via x-kubernetes-preserve-unknown-fields is not accessible in CEL expressions. This includes: - Unknown field values that are preserved by object schemas with x-kubernetes-preserve-unknown-fields. - Object properties where the property schema is of an "unknown type". An "unknown type" is recursively defined as:
	// - A schema with no type and x-kubernetes-preserve-unknown-fields set to true
	// - An array where the items schema is of an "unknown type"
	// - An object where the additionalProperties schema is of an "unknown type"
	//
	// Only property names of the form `[a-zA-Z_.-/][a-zA-Z0-9_.-/]*` are accessible. Accessible property names are escaped according to the following rules when accessed in the expression: - '__' escapes to '__underscores__' - '.' escapes to '__dot__' - '-' escapes to '__dash__' - '/' escapes to '__slash__' - Property names that exactly match a CEL RESERVED keyword escape to '__{keyword}__'. The keywords are:
	// "true", "false", "null", "in", "as", "break", "const", "continue", "else", "for", "function", "if",
	// "import", "let", "loop", "package", "namespace", "return".
	// Examples:
	// - Rule accessing a property named "namespace": {"rule": "self.__namespace__ > 0"}
	// - Rule accessing a property named "x-prop": {"rule": "self.x__dash__prop > 0"}
	// - Rule accessing a property named "redact__d": {"rule": "self.redact__underscores__d > 0"}
	//
	// Equality on arrays with x-kubernetes-list-type of 'set' or 'map' ignores element order, i.e. [1, 2] == [2, 1]. Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type:
	// - 'set': `X + Y` performs a union where the array positions of all elements in `X` are preserved and
	// non-intersecting elements in `Y` are appended, retaining their partial order.
	// - 'map': `X + Y` performs a merge where the array positions of all keys in `X` are preserved but the values
	// are overwritten by values in `Y` when the key sets of `X` and `Y` intersect. Elements in `Y` with
	// non-intersecting keys are appended, retaining their partial order.
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// Message represents the message displayed when validation fails.
	//
	// The message is required if the Rule contains line breaks. The message must not contain line breaks. If unset, the message is "failed rule: {Rule}". e.g. "must be a URL with the host matching spec.host"
	Message *string `field:"optional" json:"message" yaml:"message"`
}

