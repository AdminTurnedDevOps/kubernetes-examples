// k8s
package k8s

import (
	"time"
)

// Event is a report of an event somewhere in the cluster.
//
// It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
type KubeEventProps struct {
	// eventTime is the time when this Event was first observed.
	//
	// It is required.
	EventTime *time.Time `field:"required" json:"eventTime" yaml:"eventTime"`
	// action is what action was taken/failed regarding to the regarding object.
	//
	// It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedCount *float64 `field:"optional" json:"deprecatedCount" yaml:"deprecatedCount"`
	// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedFirstTimestamp *time.Time `field:"optional" json:"deprecatedFirstTimestamp" yaml:"deprecatedFirstTimestamp"`
	// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedLastTimestamp *time.Time `field:"optional" json:"deprecatedLastTimestamp" yaml:"deprecatedLastTimestamp"`
	// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedSource *EventSource `field:"optional" json:"deprecatedSource" yaml:"deprecatedSource"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// note is a human-readable description of the status of this operation.
	//
	// Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note *string `field:"optional" json:"note" yaml:"note"`
	// reason is why the action was taken.
	//
	// It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// regarding contains the object this Event is about.
	//
	// In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding *ObjectReference `field:"optional" json:"regarding" yaml:"regarding"`
	// related is the optional secondary object for more complex actions.
	//
	// E.g. when regarding object triggers a creation or deletion of related object.
	Related *ObjectReference `field:"optional" json:"related" yaml:"related"`
	// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
	ReportingController *string `field:"optional" json:"reportingController" yaml:"reportingController"`
	// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
	ReportingInstance *string `field:"optional" json:"reportingInstance" yaml:"reportingInstance"`
	// series is data about the Event series this event represents or nil if it's a singleton Event.
	Series *EventSeries `field:"optional" json:"series" yaml:"series"`
	// type is the type of this event (Normal, Warning), new types could be added in the future.
	//
	// It is machine-readable. This field cannot be empty for new Events.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

