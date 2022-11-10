// k8s
package k8s

import (
	"time"
)

// EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time. How often to update the EventSeries is up to the event reporters. The default event reporter in "k8s.io/client-go/tools/events/event_broadcaster.go" shows how this struct is updated on heartbeats and can guide customized reporter implementations.
type EventSeries struct {
	// count is the number of occurrences in this series up to the last heartbeat time.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// lastObservedTime is the time when last Event from the series was seen before last heartbeat.
	LastObservedTime *time.Time `field:"required" json:"lastObservedTime" yaml:"lastObservedTime"`
}

