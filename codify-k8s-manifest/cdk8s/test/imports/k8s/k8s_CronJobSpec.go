// k8s
package k8s


// CronJobSpec describes how the job execution will look like and when it will actually run.
type CronJobSpec struct {
	// Specifies the job that will be created when executing a CronJob.
	JobTemplate *JobTemplateSpec `field:"required" json:"jobTemplate" yaml:"jobTemplate"`
	// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// Specifies how to treat concurrent executions of a Job.
	//
	// Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one.
	ConcurrencyPolicy *string `field:"optional" json:"concurrencyPolicy" yaml:"concurrencyPolicy"`
	// The number of failed finished jobs to retain.
	//
	// Value must be non-negative integer. Defaults to 1.
	FailedJobsHistoryLimit *float64 `field:"optional" json:"failedJobsHistoryLimit" yaml:"failedJobsHistoryLimit"`
	// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.
	//
	// Missed jobs executions will be counted as failed ones.
	StartingDeadlineSeconds *float64 `field:"optional" json:"startingDeadlineSeconds" yaml:"startingDeadlineSeconds"`
	// The number of successful finished jobs to retain.
	//
	// Value must be non-negative integer. Defaults to 3.
	SuccessfulJobsHistoryLimit *float64 `field:"optional" json:"successfulJobsHistoryLimit" yaml:"successfulJobsHistoryLimit"`
	// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.
	//
	// Defaults to false.
	Suspend *bool `field:"optional" json:"suspend" yaml:"suspend"`
	// The time zone name for the given schedule, see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones. If not specified, this will default to the time zone of the kube-controller-manager process. The set of valid time zone names and the time zone offset is loaded from the system-wide time zone database by the API server during CronJob validation and the controller manager during execution. If no system-wide time zone database can be found a bundled version of the database is used instead. If the time zone name becomes invalid during the lifetime of a CronJob or due to a change in host configuration, the controller will stop creating new new Jobs and will create a system event with the reason UnknownTimeZone. More information can be found in https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#time-zones This is beta field and must be enabled via the `CronJobTimeZone` feature gate.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

