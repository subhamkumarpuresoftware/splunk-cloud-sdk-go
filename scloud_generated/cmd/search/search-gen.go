// Package search -- generated by scloudgen
// !! DO NOT EDIT !!
//
package search

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/scloud_generated/pkg/search"
)

// createJob -- Creates a search job.
var createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "Creates a search job.",
	RunE:  impl.CreateJob,
}

// getJob -- Return the search job with the specified search ID (SID).
var getJobCmd = &cobra.Command{
	Use:   "get-job",
	Short: "Return the search job with the specified search ID (SID).",
	RunE:  impl.GetJob,
}

// listEventsSummary -- Return events summary, for search ID (SID) search.
var listEventsSummaryCmd = &cobra.Command{
	Use:   "list-events-summary",
	Short: "Return events summary, for search ID (SID) search.",
	RunE:  impl.ListEventsSummary,
}

// listFieldsSummary -- Return fields stats summary of the events to-date, for search ID (SID) search.
var listFieldsSummaryCmd = &cobra.Command{
	Use:   "list-fields-summary",
	Short: "Return fields stats summary of the events to-date, for search ID (SID) search.",
	RunE:  impl.ListFieldsSummary,
}

// listJobs -- Return the matching list of search jobs.
var listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "Return the matching list of search jobs.",
	RunE:  impl.ListJobs,
}

// listPreviewResults -- Return the preview search results for the job with the specified search ID (SID). Can be used when a job is running to return interim results.
var listPreviewResultsCmd = &cobra.Command{
	Use:   "list-preview-results",
	Short: "Return the preview search results for the job with the specified search ID (SID). Can be used when a job is running to return interim results.",
	RunE:  impl.ListPreviewResults,
}

// listResults -- Return the search results for the job with the specified search ID (SID).
var listResultsCmd = &cobra.Command{
	Use:   "list-results",
	Short: "Return the search results for the job with the specified search ID (SID).",
	RunE:  impl.ListResults,
}

// listTimeBuckets -- Return event distribution over time of the untransformed events read to-date, for search ID(SID) search.
var listTimeBucketsCmd = &cobra.Command{
	Use:   "list-time-buckets",
	Short: "Return event distribution over time of the untransformed events read to-date, for search ID(SID) search.",
	RunE:  impl.ListTimeBuckets,
}

// updateJob -- Update the search job with the specified search ID (SID) with an action.
var updateJobCmd = &cobra.Command{
	Use:   "update-job",
	Short: "Update the search job with the specified search ID (SID) with an action.",
	RunE:  impl.UpdateJob,
}

func init() {
	searchCmd.AddCommand(createJobCmd)
	var createJobQuery string
	createJobCmd.Flags().StringVar(&createJobQuery, "query", "", "The SPL search string.")
	createJobCmd.MarkFlagRequired("query")

	var createJobAllowSideEffects string
	createJobCmd.Flags().StringVar(&createJobAllowSideEffects, "allow-side-effects", "", "Specifies whether a search that contains commands with side effects (with possible security risks) is allowed to run. type: boolean default: false")
	var createJobCollectEventSummary string
	createJobCmd.Flags().StringVar(&createJobCollectEventSummary, "collect-event-summary", "", "Specified whether a search is allowed to collect events summary during the run time.")
	var createJobCollectFieldSummary string
	createJobCmd.Flags().StringVar(&createJobCollectFieldSummary, "collect-field-summary", "", "Specified whether a search is allowed to collect Fields summary during the run time.")
	var createJobCollectTimeBuckets string
	createJobCmd.Flags().StringVar(&createJobCollectTimeBuckets, "collect-time-buckets", "", "Specified whether a search is allowed to collect Timeline Buckets summary during the run time.")
	var createJobCompletionTime string
	createJobCmd.Flags().StringVar(&createJobCompletionTime, "completion-time", "", "The time, in GMT, that the search job is finished. Empty if the search job has not completed.")
	var createJobDispatchTime string
	createJobCmd.Flags().StringVar(&createJobDispatchTime, "dispatch-time", "", "The time, in GMT, that the search job is dispatched.")
	var createJobEnablePreview string
	createJobCmd.Flags().StringVar(&createJobEnablePreview, "enable-preview", "", "Specified whether a search is allowed to collect preview results during the run time.")
	var createJobExtractAllFields string
	createJobCmd.Flags().StringVar(&createJobExtractAllFields, "extract-all-fields", "", "Specifies whether the Search service should extract all of the available fields in the data, including fields not mentioned in the SPL for the search job. Set to 'false' for better search peformance.")
	var createJobMaxTime string
	createJobCmd.Flags().StringVar(&createJobMaxTime, "max-time", "", "The number of seconds to run the search before finalizing the search. The maximum value is 21600 seconds (6 hours).")
	var createJobMessages string
	createJobCmd.Flags().StringVar(&createJobMessages, "messages", "", "")
	var createJobModule string
	createJobCmd.Flags().StringVar(&createJobModule, "module", "", "The module to run the search in. The default module is used if a module is not specified.")
	var createJobName string
	createJobCmd.Flags().StringVar(&createJobName, "name", "", "The name of the created search job.")
	var createJobPercentComplete string
	createJobCmd.Flags().StringVar(&createJobPercentComplete, "percent-complete", "", "An estimate of the percent of time remaining before the job completes.")
	var createJobPreviewAvailable string
	createJobCmd.Flags().StringVar(&createJobPreviewAvailable, "preview-available", "", "Specifies if preview results for the search job are available. The valid status values are 'unknown', 'true', and 'false'.")
	var createJobQueryParameters string
	createJobCmd.Flags().StringVar(&createJobQueryParameters, "query-parameters", "", "Represents parameters on the search job such as 'earliest' and 'latest'.")
	var createJobResolvedEarliest string
	createJobCmd.Flags().StringVar(&createJobResolvedEarliest, "resolved-earliest", "", "The earliest time speciifed as an absolute value in GMT. The time is computed based on the values you specify for the 'timezone' and 'earliest' queryParameters.")
	var createJobResolvedLatest string
	createJobCmd.Flags().StringVar(&createJobResolvedLatest, "resolved-latest", "", "The latest time specified as an absolute value in GMT. The time is computed based on the values you specify for the 'timezone' and 'earliest' queryParameters.")
	var createJobResultsAvailable string
	createJobCmd.Flags().StringVar(&createJobResultsAvailable, "results-available", "", "The number of results produced so far for the search job.")
	var createJobSid string
	createJobCmd.Flags().StringVar(&createJobSid, "sid", "", "The ID assigned to the search job.")
	var createJobStatus string
	createJobCmd.Flags().StringVar(&createJobStatus, "status", "", "status can accept values running, done, canceled, failed, ")

	searchCmd.AddCommand(getJobCmd)
	var getJobSid string
	getJobCmd.Flags().StringVar(&getJobSid, "sid", "", "The search ID.")
	getJobCmd.MarkFlagRequired("sid")

	searchCmd.AddCommand(listEventsSummaryCmd)
	var listEventsSummarySid string
	listEventsSummaryCmd.Flags().StringVar(&listEventsSummarySid, "sid", "", "The search ID.")
	listEventsSummaryCmd.MarkFlagRequired("sid")

	var listEventsSummaryCount string
	listEventsSummaryCmd.Flags().StringVar(&listEventsSummaryCount, "count", "", "The maximum number of entries to return. Set to 0 to return all available entries.")
	var listEventsSummaryEarliest string
	listEventsSummaryCmd.Flags().StringVar(&listEventsSummaryEarliest, "earliest", "", "The earliest time filter, in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%!F(MISSING)T%!T(MISSING).%!Q(MISSING)) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.")
	var listEventsSummaryField string
	listEventsSummaryCmd.Flags().StringVar(&listEventsSummaryField, "field", "", "A field to return for the result set. You can specify multiple fields of comma-separated values if multiple fields are required.")
	var listEventsSummaryLatest string
	listEventsSummaryCmd.Flags().StringVar(&listEventsSummaryLatest, "latest", "", "The latest time filter in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%!F(MISSING)T%!T(MISSING).%!Q(MISSING)) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.")
	var listEventsSummaryOffset string
	listEventsSummaryCmd.Flags().StringVar(&listEventsSummaryOffset, "offset", "", "Index of first item to return.")

	searchCmd.AddCommand(listFieldsSummaryCmd)
	var listFieldsSummarySid string
	listFieldsSummaryCmd.Flags().StringVar(&listFieldsSummarySid, "sid", "", "The search ID.")
	listFieldsSummaryCmd.MarkFlagRequired("sid")

	var listFieldsSummaryEarliest string
	listFieldsSummaryCmd.Flags().StringVar(&listFieldsSummaryEarliest, "earliest", "", "The earliest time filter, in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%!F(MISSING)T%!T(MISSING).%!Q(MISSING)) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.")
	var listFieldsSummaryLatest string
	listFieldsSummaryCmd.Flags().StringVar(&listFieldsSummaryLatest, "latest", "", "The latest time filter in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%!F(MISSING)T%!T(MISSING).%!Q(MISSING)) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.")

	searchCmd.AddCommand(listJobsCmd)

	var listJobsCount string
	listJobsCmd.Flags().StringVar(&listJobsCount, "count", "", "The maximum number of jobs that you want to return the status entries for.")
	var listJobsStatus string
	listJobsCmd.Flags().StringVar(&listJobsStatus, "status", "", "status can accept values running, done, canceled, failed, ")

	searchCmd.AddCommand(listPreviewResultsCmd)
	var listPreviewResultsSid string
	listPreviewResultsCmd.Flags().StringVar(&listPreviewResultsSid, "sid", "", "The search ID.")
	listPreviewResultsCmd.MarkFlagRequired("sid")

	var listPreviewResultsCount string
	listPreviewResultsCmd.Flags().StringVar(&listPreviewResultsCount, "count", "", "The maximum number of entries to return. Set to 0 to return all available entries.")
	var listPreviewResultsOffset string
	listPreviewResultsCmd.Flags().StringVar(&listPreviewResultsOffset, "offset", "", "Index of first item to return.")

	searchCmd.AddCommand(listResultsCmd)
	var listResultsSid string
	listResultsCmd.Flags().StringVar(&listResultsSid, "sid", "", "The search ID.")
	listResultsCmd.MarkFlagRequired("sid")

	var listResultsCount string
	listResultsCmd.Flags().StringVar(&listResultsCount, "count", "", "The maximum number of entries to return. Set to 0 to return all available entries.")
	var listResultsField string
	listResultsCmd.Flags().StringVar(&listResultsField, "field", "", "A field to return for the result set. You can specify multiple fields of comma-separated values if multiple fields are required.")
	var listResultsOffset string
	listResultsCmd.Flags().StringVar(&listResultsOffset, "offset", "", "Index of first item to return.")

	searchCmd.AddCommand(listTimeBucketsCmd)
	var listTimeBucketsSid string
	listTimeBucketsCmd.Flags().StringVar(&listTimeBucketsSid, "sid", "", "The search ID.")
	listTimeBucketsCmd.MarkFlagRequired("sid")

	searchCmd.AddCommand(updateJobCmd)
	var updateJobSid string
	updateJobCmd.Flags().StringVar(&updateJobSid, "sid", "", "The search ID.")
	updateJobCmd.MarkFlagRequired("sid")
	var updateJobStatus string
	updateJobCmd.Flags().StringVar(&updateJobStatus, "status", "", "The status to PATCH to an existing search job. The only status values you can PATCH are 'canceled' and 'finalized'. You can PATCH the 'canceled' status only to a search job that is running.")
	updateJobCmd.MarkFlagRequired("status")

}
