// Package collect -- generated by scloudgen
// !! DO NOT EDIT !!
//
package collect

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/scloud_generated/pkg/collect"
)

// createJob -- This API returns `403` if the number of collect workers is over a certain limit.
var createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "This API returns `403` if the number of collect workers is over a certain limit.",
	RunE:  impl.CreateJob,
}

// deleteJob -- Removes a job based on the job ID.
var deleteJobCmd = &cobra.Command{
	Use:   "delete-job",
	Short: "Removes a job based on the job ID.",
	RunE:  impl.DeleteJob,
}

// deleteJobs -- Removes all jobs on a tenant.
var deleteJobsCmd = &cobra.Command{
	Use:   "delete-jobs",
	Short: "Removes all jobs on a tenant.",
	RunE:  impl.DeleteJobs,
}

// getJob -- Returns a job based on the job ID.
var getJobCmd = &cobra.Command{
	Use:   "get-job",
	Short: "Returns a job based on the job ID.",
	RunE:  impl.GetJob,
}

// listJobs -- Returns a list of all jobs that belong to a tenant.
var listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "Returns a list of all jobs that belong to a tenant.",
	RunE:  impl.ListJobs,
}

// patchJob -- This API returns `403` if the number of collect workers is over a certain limit.
var patchJobCmd = &cobra.Command{
	Use:   "patch-job",
	Short: "This API returns `403` if the number of collect workers is over a certain limit.",
	RunE:  impl.PatchJob,
}

// patchJobs -- This is a non-atomic operation and the results are returned as a list with each job patch result as its element. This API returns `200 OK` regardless of how many jobs were successfully patched. You must read the response body to find out if all jobs are patched. When the API is called, the `jobIDs` or `connectorID` must be specified. Do not specify more than one of them at the same time. This API returns `403` if the number of collect workers is over a certain limit.
var patchJobsCmd = &cobra.Command{
	Use:   "patch-jobs",
	Short: "This is a non-atomic operation and the results are returned as a list with each job patch result as its element. This API returns `200 OK` regardless of how many jobs were successfully patched. You must read the response body to find out if all jobs are patched. When the API is called, the `jobIDs` or `connectorID` must be specified. Do not specify more than one of them at the same time. This API returns `403` if the number of collect workers is over a certain limit.",
	RunE:  impl.PatchJobs,
}

func init() {
	collectCmd.AddCommand(createJobCmd)

	var createJobAuthorization string
	createJobCmd.Flags().StringVar(&createJobAuthorization, "authorization", "", "This is a required parameter.The access token provided by the user (obtained from a known identity provider).")
	createJobCmd.MarkFlagRequired("authorization")

	var createJobConnectorID string
	createJobCmd.Flags().StringVar(&createJobConnectorID, "connector-ID", "", "This is a required parameter.This is a required parameter.The ID of the connector used in the job.")
	createJobCmd.MarkFlagRequired("connector-ID")

	var createJobName string
	createJobCmd.Flags().StringVar(&createJobName, "name", "", "This is a required parameter.This is a required parameter.")
	createJobCmd.MarkFlagRequired("name")

	var createJobParameters string
	createJobCmd.Flags().StringVar(&createJobParameters, "parameters", "", "This is a required parameter.This is a required parameter.The configuration of the connector used in the job.")
	createJobCmd.MarkFlagRequired("parameters")

	var createJobScalePolicy string
	createJobCmd.Flags().StringVar(&createJobScalePolicy, "scale-policy", "", "This is a required parameter.This is a required parameter.")
	createJobCmd.MarkFlagRequired("scale-policy")

	var createJobSchedule string
	createJobCmd.Flags().StringVar(&createJobSchedule, "schedule", "", "This is a required parameter.This is a required parameter.The cron schedule, in UTC time format.")
	createJobCmd.MarkFlagRequired("schedule")

	var createJobEventExtraFields string
	createJobCmd.Flags().StringVar(&createJobEventExtraFields, "event-extra-fields", "", "")

	var createJobScheduled string
	createJobCmd.Flags().StringVar(&createJobScheduled, "scheduled", "false", "Defines whether a job is scheduled or not")

	collectCmd.AddCommand(deleteJobCmd)

	var deleteJobAuthorization string
	deleteJobCmd.Flags().StringVar(&deleteJobAuthorization, "authorization", "", "This is a required parameter.The access token provided by the user (obtained from a known identity provider).")
	deleteJobCmd.MarkFlagRequired("authorization")

	var deleteJobJobId string
	deleteJobCmd.Flags().StringVar(&deleteJobJobId, "job-id", "", "This is a required parameter.The job ID.")
	deleteJobCmd.MarkFlagRequired("job-id")

	collectCmd.AddCommand(deleteJobsCmd)

	var deleteJobsAuthorization string
	deleteJobsCmd.Flags().StringVar(&deleteJobsAuthorization, "authorization", "", "This is a required parameter.The access token provided by the user (obtained from a known identity provider).")
	deleteJobsCmd.MarkFlagRequired("authorization")

	collectCmd.AddCommand(getJobCmd)

	var getJobAuthorization string
	getJobCmd.Flags().StringVar(&getJobAuthorization, "authorization", "", "This is a required parameter.The access token provided by the user (obtained from a known identity provider).")
	getJobCmd.MarkFlagRequired("authorization")

	var getJobJobId string
	getJobCmd.Flags().StringVar(&getJobJobId, "job-id", "", "This is a required parameter.The job ID.")
	getJobCmd.MarkFlagRequired("job-id")

	collectCmd.AddCommand(listJobsCmd)

	var listJobsAuthorization string
	listJobsCmd.Flags().StringVar(&listJobsAuthorization, "authorization", "", "This is a required parameter.The access token provided by the user (obtained from a known identity provider).")
	listJobsCmd.MarkFlagRequired("authorization")

	var listJobsConnectorId string
	listJobsCmd.Flags().StringVar(&listJobsConnectorId, "connector-id", "", "Specifies the connector ID used to filter jobs. A tailing wildcard is supported for the connector ID tag. If no wildcard is used then an exact match is used. Examples: * `my-connector:v1.0.0` selects `my-connector` connector with an exact match with tag v1.0.0 * `my-connector` selects `my-connector` connector with an exact match. Note as no tag is specified it actually refers to latest. * `my-connector:v1.*` selects all `my-connector` connectors with tags starting with v1., e.g. v1.0, v1.1.1, v1.2-alpha, etc. * `my-connector:*` selects all `my-connector` connectors with any tag.")

	collectCmd.AddCommand(patchJobCmd)

	var patchJobAuthorization string
	patchJobCmd.Flags().StringVar(&patchJobAuthorization, "authorization", "", "This is a required parameter.The access token provided by the user (obtained from a known identity provider).")
	patchJobCmd.MarkFlagRequired("authorization")

	var patchJobJobId string
	patchJobCmd.Flags().StringVar(&patchJobJobId, "job-id", "", "This is a required parameter.The job ID.")
	patchJobCmd.MarkFlagRequired("job-id")

	var patchJobConnectorID string
	patchJobCmd.Flags().StringVar(&patchJobConnectorID, "connector-ID", "", "The ID of the connector used in the job.")

	var patchJobEventExtraFields string
	patchJobCmd.Flags().StringVar(&patchJobEventExtraFields, "event-extra-fields", "", "")

	var patchJobName string
	patchJobCmd.Flags().StringVar(&patchJobName, "name", "", "The job name")

	var patchJobParameters string
	patchJobCmd.Flags().StringVar(&patchJobParameters, "parameters", "", "The configuration of the connector used in the job.")

	var patchJobScalePolicy string
	patchJobCmd.Flags().StringVar(&patchJobScalePolicy, "scale-policy", "", "")

	var patchJobSchedule string
	patchJobCmd.Flags().StringVar(&patchJobSchedule, "schedule", "", "The cron schedule, in UTC time format.")

	var patchJobScheduled string
	patchJobCmd.Flags().StringVar(&patchJobScheduled, "scheduled", "false", "Defines wheather a job is scheduled or not")

	collectCmd.AddCommand(patchJobsCmd)

	var patchJobsAuthorization string
	patchJobsCmd.Flags().StringVar(&patchJobsAuthorization, "authorization", "", "This is a required parameter.The access token provided by the user (obtained from a known identity provider).")
	patchJobsCmd.MarkFlagRequired("authorization")

	var patchJobsConnectorID string
	patchJobsCmd.Flags().StringVar(&patchJobsConnectorID, "connector-ID", "", "The ID of the connector used in the job.")

	var patchJobsConnectorId string
	patchJobsCmd.Flags().StringVar(&patchJobsConnectorId, "connector-id", "", "Specifies the connector ID used to filter jobs. A tailing wildcard is supported for the connector ID tag. If no wildcard is used then an exact match is used. Examples: * `my-connector:v1.0.0` selects `my-connector` connector with an exact match with tag v1.0.0 * `my-connector` selects `my-connector` connector with an exact match. Note as no tag is specified it actually refers to latest. * `my-connector:v1.*` selects all `my-connector` connectors with tags starting with v1., e.g. v1.0, v1.1.1, v1.2-alpha, etc. * `my-connector:*` selects all `my-connector` connectors with any tag.")

	var patchJobsEventExtraFields string
	patchJobsCmd.Flags().StringVar(&patchJobsEventExtraFields, "event-extra-fields", "", "")

	var patchJobsJobIDs []string
	patchJobsCmd.Flags().StringSliceVar(&patchJobsJobIDs, "job-ids", nil, "The job ID list.")

	var patchJobsScalePolicy string
	patchJobsCmd.Flags().StringVar(&patchJobsScalePolicy, "scale-policy", "", "")

}
